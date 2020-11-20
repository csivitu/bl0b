package mux

import (
	"log"
	"strconv"
	"time"

	"github.com/csivitu/bl0b/ctftime"
	"github.com/csivitu/bl0b/db"
	"github.com/csivitu/bl0b/utils"

	"github.com/bwmarrin/discordgo"
)

// UpcomingEvents returns 3 upcoming events from the database
func (m *Mux) UpcomingEvents(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	DB := db.New()
	defer DB.Close()

	numberOfEvents := 3
	events, err := DB.GetEventsByStatus(utils.Upcoming)

	events = events[:numberOfEvents]
	if err != nil {
		log.Println(err)
	}

	message := ":triangular_flag_on_post: **_CTFs This Week_** :triangular_flag_on_post:\n\n"
	message = attachEventData(message, &events)

	ds.ChannelMessageSend(dm.ChannelID, message)
}

// OngoingEvents returns ongoing events from the database
func (m *Mux) OngoingEvents(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	DB := db.New()
	defer DB.Close()

	events, err := DB.GetEventsByStatus(utils.Ongoing)
	if err != nil {
		log.Println(err)
	}

	message := ":triangular_flag_on_post: **_Ongoing CTFs_** :triangular_flag_on_post:\n\n"
	message = attachEventData(message, &events)

	ds.ChannelMessageSend(dm.ChannelID, message)
}

func attachEventData(message string, events *ctftime.Events) string {
	for i := 0; i < len(*events); i++ {
		event := (*events)[i]
		weight := strconv.FormatFloat(event.Weight, 'f', 2, 64)

		message += "**" + event.Title + "**\n"

		message += "Organizers:\n"

		// TODO: Temporary hack - return only the first Organizer's name
		// for j := 0; j < len(event.Organizers); j++ {
		// 	message += strconv.Itoa(j+1) + ". **" + event.Organizers[j].Name + "**\n"
		// }

		message += "1. **" + event.Organizer + "**\n"
		message += "Weight: **" + weight + "**\n"
		message += "Official URL: " + event.URL + "\n"
		message += "CTFtime URL: " + event.CtftimeURL + "\n"
		message += "Format: " + event.Format + "\n"
		message += "Starts at: " + event.Start.Format(time.RFC1123) + "\n"
		message += "Ends at: " + event.Finish.Format(time.RFC1123) + "\n"
		message += "\n"
	}

	if len(*events) == 0 {
		message = "0 CTFs found :slight_frown:"
	}

	return message
}
