package routines

import (
	"log"
	"time"

	"github.com/csivitu/bl0b/ctftime"

	"github.com/csivitu/bl0b/db"
)

func computeStatus(start time.Time, end time.Time) string {
	t := time.Now().Unix()
	startUnix := start.Unix()
	endUnix := end.Unix()

	if startUnix <= t && t < endUnix {
		return "ongoing"
	}

	if t < startUnix && t < endUnix {
		return "upcoming"
	}

	return "over"
}

func handleOngoing(event *ctftime.Event, DB *db.Database) {
	// If the status is already ongoing, return
	if event.Status == "ongoing" {
		return
	}

	err := DB.ModifyEventStatus(event.ID, "ongoing")
	if err != nil {
		log.Println("Error while modifying event " + event.Title)
		log.Println(err)
	}
}

func handleOver(event *ctftime.Event, DB *db.Database) {
	err := DB.DeleteEventByID(event.ID)
	if err != nil {
		log.Println("Error while deleting event " + event.Title)
		log.Println(err)
	}
}

func analyze() {
	DB := db.New()
	defer DB.Close()

	events, err := DB.GetEvents()
	if err != nil {
		log.Println("Analyze could not get events from database!")
		panic(err)
	}

	for _, event := range events {
		status := computeStatus(event.Start, event.Finish)

		switch status {
		case "ongoing":
			{
				handleOngoing(&event, DB)
			}
		case "over":
			{
				handleOver(&event, DB)
			}
		}
	}
}

// Analyze rows in the database to change status
// from `upcoming` to `ongoing`
func Analyze(t time.Duration) {
	ticker := time.NewTicker(t)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				analyze()
			}
		}
	}()

	log.Println("Analyzing the DB to update the status of each event!")
}
