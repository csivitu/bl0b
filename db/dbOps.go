package db

import (
	"github.com/csivitu/bl0b/ctftime"
)

// AddEvent adds an event in the database
func (DB *Database) AddEvent(event *ctftime.Event) error {
	queryString := `
		INSERT INTO events (
			ID, CtfID, FormatID, Logo,
			PublicVotable, LiveFeed,
			Location, CtftimeURL,
			Participants, Start,
			Format, Restrictions,
			IsVotableNow, URL, Title,
			Weight, Description,
			Finish, OnSite
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`
	
	_, err := DB.db.Exec(queryString,
		event.ID, event.CtfID, event.FormatID, event.Logo,
		event.PublicVotable, event.LiveFeed,
		event.Location, event.CtftimeURL,
		event.Participants, event.Start,
		event.Format, event.Restrictions,
		event.IsVotableNow, event.URL, event.Title,
		event.Weight, event.Description,
		event.Finish, event.OnSite,
	)

	return err
}
