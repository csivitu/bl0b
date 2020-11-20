package routines

import (
	"log"
	"time"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/db"
	"github.com/csivitu/bl0b/utils"
)

func populate(t time.Time) {
	log.Println("Populating database with upcoming CTFs!")
	DB := db.New()
	defer DB.Close()

	ctf := ctftime.New()

	duration := time.Hour*24*7

	// 7 days before and 7 days after
	events, err := ctf.GetEvents(10, t.Add(-duration), t.Add(duration))

	if err != nil {
		log.Println(err)
	}

	err = DB.AddEvents(&events)
	if err != nil {
		log.Println(err)
	}
}

// Populate run every interval and adds items to the database
func Populate(t time.Duration) {
	// Instantly populate once in a goroutine
	go populate(time.Now())
	utils.SetInterval(populate, t)

	log.Println("The database will be populated periodically.")
}
