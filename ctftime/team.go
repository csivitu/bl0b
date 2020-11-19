package ctftime

// Team defines the structure of a single team on CTFtime
type Team struct {
	Name      string      `json:"name"`
	Country   string      `json:"country"`
	ID        int         `json:"id"`
	Aliases   []string    `json:"aliases"`
	// Ratings   []Rating    `json:"rating"`
}

// type Rating struct {

// }
