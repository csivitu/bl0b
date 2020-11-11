package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Session is declared in a global namespace
var Session, _ = discordgo.New()

func init() {
	Session.Token = os.Getenv("DG_TOKEN")

	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
}

func main() {
	fmt.Println(`
	┌───────────┐
	│▛▀▖▜    ▌  │
	│▙▄▘▐ ▞▀▖▛▀▖│
	│▌ ▌▐ ▌ ▌▌ ▌│
	│▀▀  ▘▝▀ ▀▀ │
	└───────────┘
	`)

	flag.Parse()

	if Session.Token == "" {
		log.Println("Discord Authentication Token not provided.")
		return
	}

	log.Println("Blob is running, press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}
