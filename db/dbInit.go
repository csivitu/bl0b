package db

import (
	"fmt"
	"strconv"
	"log"
	"os"
)

var (
	dbUser string
	dbPass string
	dbIP string
	dbPort int
	dbName string
)

func init() {
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")

	if dbUser == "" || dbPass == "" {
		log.Fatal("Environment variables DB_USER and DB_PASS were not set.")
	}

	dbIP = os.Getenv("DB_IP")
	dbPortStr := os.Getenv("DB_PORT")

	if dbIP == "" {
		log.Println("DB_IP was not set, using 127.0.0.1.")
		dbIP = "127.0.0.1"
	}

	if dbPortStr == "" {
		log.Println("DB_PORT was not set, using 3306.")
		dbPort = 3306
	} else {
		var err error
		dbPort, err = strconv.Atoi(dbPortStr)

		if err != nil {
			log.Fatal("Invalid DB_PORT.")
		}
	}

	dbName = "bl0b"
}

// Init is used to initialize the SQL Database
func Init() {
	db := dbConn(dbUser, dbPass, dbIP, dbPort, "")

	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		panic(err)
	}

	db.Close()

	db = dbConn(dbUser, dbPass, dbIP, dbPort, dbName)

	_, err = db.Exec("DROP TABLE IF EXISTS events")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE events (
			ID            INT NOT NULL,
			CtfID         INT NOT NULL,
			FormatID      INT NOT NULL,
			Logo          VARCHAR(50),
			PublicVotable BOOL,
			LiveFeed      VARCHAR(50),
			Location      VARCHAR(50),
			CtftimeURL    VARCHAR(200) NOT NULL,
			Participants  INT NOT NULL,
			Start         DATETIME,
			Format        VARCHAR(50) NOT NULL,
			Restrictions  VARCHAR(100) NOT NULL,
			IsVotableNow  BOOL,
			URL           VARCHAR(200) NOT NULL,
			Title         VARCHAR(100) NOT NULL,
			Weight        DOUBLE NOT NULL,
			Description   VARCHAR(300) NOT NULL,
			Finish        DATETIME,
			OnSite        BOOL
		)
	`)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Println("Database initialized, tables created.")
}
