package ctftime

import (
	"time"
)

// Event defines a single event obtained from CTFtime
type Event struct {
	Organizers    []Organizer `json:"organizers"`
	OnSite        bool        `json:"onsite"`
	Finish        time.Time   `json:"finish"`
	Description   string      `json:"description"`
	Weight        float64     `json:"weight"`
	Title         string      `json:"title"`
	URL           string      `json:"url"`
	IsVotableNow  bool        `json:"is_votable_now"`
	Restrictions  string      `json:"restrictions"`
	Format        string      `json:"format"`
	Start         time.Time   `json:"start"`
	Participants  int         `json:"participants"`
	CtftimeURL    string      `json:"ctftime_url"`
	Location      string      `json:"location"`
	LiveFeed      string      `json:"live_feed"`
	PublicVotable bool        `json:"public_votable"`
	Duration      Duration    `json:"duration"`
	Logo          string      `json:"logo"`
	FormatID      int         `json:"format_id"`
	ID            int         `json:"id"`
	CtfID         int         `json:"ctf_id"`
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
