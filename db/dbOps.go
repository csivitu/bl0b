package db

import (
	"github.com/jmoiron/sqlx"
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

	queryString := fmt.Sprintf("INSERT INTO events (%s) VALUES (%s)", params, values)

	tx := DB.db.MustBegin()
	for _, event := range *events {
		tx.NamedExec(queryString, event)
	}

	err := tx.Commit()
	return err
}

// GetEvents returns Events from the database
func (DB *Database) GetEvents() (ctftime.Events, error) {
	rows, err := DB.db.Queryx("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return DB.getEventsFromRows(rows)
}

// GetEventsByStatus returns events depending upon the Status attribute
func (DB *Database) GetEventsByStatus(status string) (ctftime.Events, error) {
	rows, err := DB.db.Queryx("SELECT * FROM events WHERE Status=$1", status)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return DB.getEventsFromRows(rows)
}


func (DB *Database) getEventsFromRows(rows *sqlx.Rows) (ctftime.Events, error) {
	var (
		events ctftime.Events
		err error
	)

	for rows.Next() {
		event := ctftime.Event{}

		err = rows.StructScan(&event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, err
}
