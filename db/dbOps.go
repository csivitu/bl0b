package db

import (
	"fmt"
	"strings"

	"github.com/csivitu/bl0b/ctftime"
)

// AddEvent adds an event in the database
func (DB *Database) AddEvent(event *ctftime.Event) error {
	params := `
		ID, CtfID, FormatID, Logo,
		PublicVotable, LiveFeed,
		Location, CtftimeURL,
		Participants, Start,
		Format, Restrictions,
		IsVotableNow, URL, Title,
		Weight, Description,
		Finish, OnSite
	`

	params = strings.ReplaceAll(params, "\t", "")
	params = strings.ReplaceAll(params, "\n", " ")

	values := ":" + strings.ReplaceAll(params, ", ", ", :")[1:]

	queryString := fmt.Sprintf("INSERT INTO events (%s) VALUES (%s)", params, values)

	_, err := DB.db.NamedExec(queryString, event)

	return err
}

// GetEvents returns Events from the database
func (DB *Database) GetEvents() (*ctftime.Events, error) {
	rows, err := DB.db.Queryx("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events ctftime.Events

	for rows.Next() {
		event := ctftime.Event{}

		err = rows.StructScan(&event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return &events, err
}
