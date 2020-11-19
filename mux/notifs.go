package mux

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/csivitu/bl0b/db"
)

// NotifRegister is used to register for notifications
func (m *Mux) NotifRegister(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	DB := db.New()
	defer DB.Close()

	err := DB.RegisterForNotifs(dm.ChannelID)
	if err != nil {
		log.Println("Could not register " + dm.ChannelID)

		if strings.Contains(err.Error(), "Duplicate") {
			ds.ChannelMessageSend(dm.ChannelID, "You're already registered :love_you_gesture:")
			return
		}
		ds.ChannelMessageSend(dm.ChannelID, "Registration failed :cry:")
		return
	}

	ds.ChannelMessageSend(dm.ChannelID, "Successfully registered for notifications! :wine_glass:")
}
