package routines

import (
	"log"
	"time"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/db"
)

// Populate run every interval and adds items to the database
func Populate() {
	ticker := time.NewTicker(15 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				{
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
			}
		}
	}()

	log.Println("The database will be populated periodically.")
}
