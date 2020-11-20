package db

import (
	"fmt"
	"strings"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/utils"
)

// AddEvents adds a slice of Events in the database
func (DB *Database) AddEvents(events *ctftime.Events) error {
	params := `
		ID, CtfID, FormatID, Logo,
		PublicVotable, LiveFeed,
		Location, CtftimeURL,
		Participants, Start,
		Format, Restrictions,
		IsVotableNow, URL, Title,
		Weight, Description,
		Finish, OnSite, Organizer,
		Status
	`

	params = strings.ReplaceAll(params, "\t", "")
	params = strings.ReplaceAll(params, "\n", " ")

	values := ":" + strings.ReplaceAll(params, ", ", ", :")[1:]

	updates := strings.Split(params, ", ")

	update := ""

	for _, u := range updates {
		update += fmt.Sprintf("%s=VALUES(%s), ", u, u)
	}

	update = update[:len(update)-2]

	queryString := fmt.Sprintf("INSERT INTO events (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s", params, values, update)

	tx := DB.db.MustBegin()
	for _, event := range *events {
		event.Status = utils.ComputeStatus(event.Start, event.Finish)
		_, err := tx.NamedExec(queryString, event)
		if err != nil {
			fmt.Println(err)
		}
	}

	err := tx.Commit()
	return err
}

// GetEvents returns Events from the database
func (DB *Database) GetEvents() (ctftime.Events, error) {
	var events ctftime.Events
	err := DB.db.Select(&events, "SELECT * FROM events ORDER BY Start")

	return events, err
}

// GetEventsByStatus returns events depending upon the Status attribute
func (DB *Database) GetEventsByStatus(status utils.Status) (ctftime.Events, error) {
	var events ctftime.Events
	err := DB.db.Select(&events, "SELECT * FROM events WHERE Status=? ORDER BY Start", status)

	return events, err
}

// ModifyEventStatus modifies the status of the event
// identified by the eventID
func (DB *Database) ModifyEventStatus(eventID int, status utils.Status) error {
	queryString := "UPDATE events SET Status=:status WHERE ID=:id"

	_, err := DB.db.NamedExec(queryString,
		map[string]interface{}{
			"id":     eventID,
			"status": status,
		})

	return err
}

// DeleteEventByID deletes the event with ID eventID
func (DB *Database) DeleteEventByID(eventID int) error {
	queryString := "DELETE FROM events WHERE ID=:id"

	_, err := DB.db.NamedExec(queryString,
		map[string]interface{}{
			"id": eventID,
		})

	return err
}
