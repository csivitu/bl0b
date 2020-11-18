package notifs

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/csivitu/bl0b/db"
)

// NotifHandler objects will be used to notify channels
type NotifHandler struct {
	Session *discordgo.Session
}

// NewNotifHandler creates a new instance of Notifier
func NewNotifHandler(Session *discordgo.Session) *NotifHandler {
	return &NotifHandler{
		Session: Session,
	}
}

// Notify sends notifications to ChannelID with message
func (n *NotifHandler) Notify(ChannelID string, message string) {
	n.Session.ChannelMessageSend(ChannelID, message)
}

// NotifyAll calls Notify for all channels in the database
func (n *NotifHandler) NotifyAll(message string) {
	DB := db.New()
	channels, err := DB.GetRegisteredChannels()
	if err != nil {
		log.Println(err)
		log.Println("Error in NotifyAll")
	}

	for _, channel := range channels {
		n.Notify(channel, message)
	}
}
