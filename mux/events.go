package mux

import (
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/csivitu/bl0b/ctftime"
)

// UpcomingEvents returns upcoming events in the next 5 days
func (m *Mux) UpcomingEvents(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	ctf := ctftime.New()

	now := time.Now().Unix() * 1000

	events, err := ctf.GetEvents(3, now, now+1000*60*60*24*7)

	if err != nil {
		log.Println(err)
	}

	message := ":triangular_flag_on_post: **_CTFs This Week_** :triangular_flag_on_post:\n\n"

	for i := 0; i < len(events); i++ {
		event := events[i]
		weight := strconv.FormatFloat(event.Weight, 'f', 2, 64)
		
		message += "**" + event.Title + "**\n"

		message += "Organizers:\n"
		for j := 0; j < len(event.Organizers); j++ {
			message += strconv.Itoa(j + 1) + ". **" + event.Organizers[0].Name + "**\n"
		}
		message += "Weight: **" + weight + "**\n"
		message += "Official URL: " + event.URL + "\n"
		message += "CTFtime URL: " + event.CtftimeURL + "\n"
		message += "Format: " + event.Format + "\n"
		message += "Starts at: " + event.Start + "\n"
		message += "Ends at: " + event.Finish + "\n"
		message += "\n"
	}

	ds.ChannelMessageSend(dm.ChannelID, message)
}
