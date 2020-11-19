package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/csivitu/bl0b/routines"

	_ "github.com/go-sql-driver/mysql"

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
___.   .__  _______ ___.    
\_ |__ |  | \   _  \\_ |__  
 | __ \|  | /  /_\  \| __ \ 
 | \_\ \  |_\  \_/   \ \_\ \
 |___  /____/\_____  /___  /
     \/            \/    \/
 	`)

	flag.Parse()

	if Session.Token == "" {
		log.Println("Discord Authentication Token not provided.")
		return
	}

	err := Session.Open()
	if err != nil {
		log.Printf("Error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	routines.Populate(time.Hour)
	routines.Analyze(time.Minute)

	log.Println("bl0b is running, press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}
