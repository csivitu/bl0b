package routines

import (
	"fmt"
	"log"
	"time"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/db"
	"github.com/csivitu/bl0b/notifs"
	"github.com/csivitu/bl0b/utils"
)

func handleUpcoming(event *ctftime.Event, DB *db.Database) {
	if event.Status == utils.Upcoming {
		return
	}

	err := DB.ModifyEventStatus(event.ID, utils.Upcoming)
	if err != nil {
		log.Println("Error while modifying event " + event.Title + " in handleUpcoming")
	}
}

func handleOngoing(event *ctftime.Event, DB *db.Database, n *notifs.NotifHandler) {
	// If the status is already ongoing, return
	if event.Status == utils.Ongoing {
		return
	}

	err := DB.ModifyEventStatus(event.ID, utils.Ongoing)
	if err != nil {
		log.Println("Error while modifying event " + event.Title + " in handleOngoing")
		log.Println(err)
	}

	n.NotifyAll(fmt.Sprintf("@here %s is now live! Head to %s :triangular_flag_on_post:", event.Title, event.URL))
}

func handleOver(event *ctftime.Event, DB *db.Database, n *notifs.NotifHandler) {
	err := DB.DeleteEventByID(event.ID)
	if err != nil {
		log.Println("Error while deleting event " + event.Title)
		log.Println(err)
	}

	if event.Status != utils.Over {
		n.NotifyAll(fmt.Sprintf("%s is over, scoreboard will be available here: %s :partying_face:", event.Title, event.CtftimeURL))
	}
}

func analyze(n *notifs.NotifHandler) {
	DB := db.New()
	defer DB.Close()

	events, err := DB.GetEvents()
	if err != nil {
		log.Println("Analyze could not get events from database!")
		panic(err)
	}

	for _, event := range events {
		status := utils.ComputeStatus(event.Start, event.Finish)

		switch status {
		case utils.Upcoming:
			handleUpcoming(&event, DB)
		case utils.Ongoing:
			handleOngoing(&event, DB, n)
		case utils.Over:
			handleOver(&event, DB, n)
		}
	}
}

// Analyze rows in the database to change status
// from `upcoming` to `ongoing`
func Analyze(t time.Duration, n *notifs.NotifHandler) {
	utils.SetInterval(
		func(_ time.Time) {
			analyze(n)
		},
		t,
	)

	log.Println("Analyzing the DB to update the status of each event!")
}
