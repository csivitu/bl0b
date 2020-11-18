package db

import (
	"fmt"
	"strings"

	"github.com/csivitu/bl0b/ctftime"
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
		Finish, OnSite, Organizer
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
		tx.NamedExec(queryString, event)
	}

	err := tx.Commit()
	return err
}

// GetEvents returns Events from the database
func (DB *Database) GetEvents() (ctftime.Events, error) {
	var events ctftime.Events
	err := DB.db.Select(&events, "SELECT * FROM events")

	return events, err
}

// GetEventsByStatus returns events depending upon the Status attribute
func (DB *Database) GetEventsByStatus(status string) (ctftime.Events, error) {
	var events ctftime.Events
	err := DB.db.Select(&events, "SELECT * FROM events WHERE Status=?", status)

	return events, err
}

// ModifyEventStatus modifies the status of the event
// identified by the eventID
func (DB *Database) ModifyEventStatus(eventID int, status string) error {
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
