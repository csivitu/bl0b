package routines

import (
	"log"
	"time"

	"github.com/csivitu/bl0b/utils"

	"github.com/csivitu/bl0b/ctftime"

	"github.com/csivitu/bl0b/db"
)

func computeStatus(start time.Time, finish time.Time) string {
	t := time.Now()

	if t.After(start) && t.Before(finish) {
		return "ongoing"
	}

	if t.Before(start) && t.Before(finish) {
		return "upcoming"
	}

	return "over"
}

func handleUpcoming(event *ctftime.Event, DB *db.Database) {
	if event.Status == "upcoming" {
		return
	}

	err := DB.ModifyEventStatus(event.ID, "upcoming")
	if err != nil {
		log.Println("Error while modifying event " + event.Title + " in handleUpcoming")
	}
}

func handleOngoing(event *ctftime.Event, DB *db.Database) {
	// If the status is already ongoing, return
	if event.Status == "ongoing" {
		return
	}

	err := DB.ModifyEventStatus(event.ID, "ongoing")
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
		case "upcoming":
			handleUpcoming(&event, DB)
		case "ongoing":
			handleOngoing(&event, DB)
		case "over":
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
