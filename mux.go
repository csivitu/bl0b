package main

import (
	"github.com/csivitu/bl0b/mux"
)

// Router is registered as a global variable to allow easy access to the
// multiplexer throughout the bot.
var Router = mux.New()

func init() {
	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.

	Session.AddHandler(Router.OnMessageCreate)

	// Register the build-in help command.
	Router.Route("help", "Display this message.", Router.Help)
	Router.Route("upcoming", "Shows next 3 upcoming CTFs this week", Router.Events)
}
