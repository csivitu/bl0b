package ctftime

import (
	"time"
)

// Event defines a single event obtained from CTFtime
type Event struct {
	Organizers    []Organizer `json:"organizers"`
	OnSite        bool        `json:"onsite" db:"OnSite"`
	Finish        time.Time   `json:"finish" db:"Finish"`
	Description   string      `json:"description" db:"Description"`
	Weight        float64     `json:"weight" db:"Weight"`
	Title         string      `json:"title" db:"Title"`
	URL           string      `json:"url" db:"URL"`
	IsVotableNow  bool        `json:"is_votable_now" db:"IsVotableNow"`
	Restrictions  string      `json:"restrictions" db:"Restrictions"`
	Format        string      `json:"format" db:"Format"`
	Start         time.Time   `json:"start" db:"Start"`
	Participants  int         `json:"participants" db:"Participants"`
	CtftimeURL    string      `json:"ctftime_url" db:"CtftimeURL"`
	Location      string      `json:"location" db:"Location"`
	LiveFeed      string      `json:"live_feed" db:"LiveFeed" db:"LiveFeed"`
	PublicVotable bool        `json:"public_votable" db:"PublicVotable"`
	Duration      Duration    `json:"duration"`
	Logo          string      `json:"logo" db:"Logo"`
	FormatID      int         `json:"format_id" db:"FormatID"`
	ID            int         `json:"id" db:"ID"`
	CtfID         int         `json:"ctf_id" db:"CtfID"`
	Status        string      `db:"Status"`
	Organizer     string      `db:"Organizer"`
	// TODO: Later on replace Organizer with list of Organizers in new table
}

// Organizer defines the structure of an organizer as obtained
// from CTFtime
type Organizer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Duration consists of the number of hours and number of days
// for a CTFtime event
type Duration struct {
	Hours int `json:"hours"`
	Days  int `json:"days"`
}

// Events is a slice of type Event
type Events []Event
