package routines

import (
	"log"
	"time"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/db"
	"github.com/csivitu/bl0b/utils"
)

func computeStatus(start time.Time, finish time.Time) ctftime.Status {
	t := time.Now()

	if t.After(start) && t.Before(finish) {
		return ctftime.Ongoing
	}

	if t.Before(start) && t.Before(finish) {
		return ctftime.Upcoming
	}

	return ctftime.Over
}

func handleUpcoming(event *ctftime.Event, DB *db.Database) {
	if event.Status == ctftime.Upcoming {
		return
	}

	err := DB.ModifyEventStatus(event.ID, ctftime.Upcoming)
	if err != nil {
		log.Println("Error while modifying event " + event.Title + " in handleUpcoming")
	}
}

func handleOngoing(event *ctftime.Event, DB *db.Database) {
	// If the status is already ongoing, return
	if event.Status == ctftime.Ongoing {
		return
	}

	err := DB.ModifyEventStatus(event.ID, ctftime.Ongoing)
	if err != nil {
		log.Println("Error while modifying event " + event.Title + " in handleOngoing")
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
		case ctftime.Upcoming:
			handleUpcoming(&event, DB)
		case ctftime.Ongoing:
			handleOngoing(&event, DB)
		case ctftime.Over:
			handleOver(&event, DB)
		}
	}
}

// Analyze rows in the database to change status
// from `upcoming` to `ongoing`
func Analyze(t time.Duration) {
	utils.SetInterval(
		func(_ time.Time) {
			analyze()
		},
		t,
	)

	log.Println("Analyzing the DB to update the status of each event!")
}
