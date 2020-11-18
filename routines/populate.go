package routines

import (
	"log"
	"time"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/db"
)

func populate(t time.Time) {
	log.Println("Populating database with upcoming CTFs!")
	DB := db.New()
	defer DB.Close()

	ctf := ctftime.New()

	now := t.Unix()
	events, err := ctf.GetEvents(10, now, now+60*60*24*7)

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
	ticker := time.NewTicker(t)

	done := make(chan bool)

	go func() {
		populate(time.Now())
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				populate(t)
			}
		}
	}()

	log.Println("The database will be populated periodically.")
}
