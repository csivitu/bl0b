package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var (
	dbUser string
	dbPass string
	dbIP   string
	dbPort int
	dbName string
)

// Database is a wrapper around the SQL database
// to allow addition of methods
type Database struct {
	db *sqlx.DB
}

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
func Init() error {
	db := dbConn(dbUser, dbPass, dbIP, dbPort, "")

	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return err
	}

	db.Close()

	DB := New()
	log.Println("Database connection successful!")

	defer DB.Close()

	_, err = DB.db.Exec("DROP TABLE IF EXISTS events")
	if err != nil {
		return err
	}

	// TODO: Add Organizers and Duration
	_, err = DB.db.Exec(`
		CREATE TABLE events (
			ID            INT NOT NULL PRIMARY KEY,
			CtfID         INT NOT NULL,
			FormatID      INT NOT NULL,
			Logo          VARCHAR(100),
			PublicVotable BOOL,
			LiveFeed      VARCHAR(100),
			Location      VARCHAR(100),
			CtftimeURL    VARCHAR(200) NOT NULL,
			Participants  INT NOT NULL,
			Start         DATETIME,
			Format        VARCHAR(50) NOT NULL,
			Restrictions  VARCHAR(100) NOT NULL,
			IsVotableNow  BOOL,
			URL           VARCHAR(200) NOT NULL,
			Title         VARCHAR(100) NOT NULL,
			Weight        DOUBLE NOT NULL,
			Description   VARCHAR(2000) NOT NULL,
			Finish        DATETIME,
			OnSite        BOOL,
			Status        ENUM('ongoing', 'upcoming', 'over'),
			Organizer     VARCHAR(100)
		)
	`)
	if err != nil {
		return err
	}

	_, err = DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS notify (
			ChannelID    VARCHAR(30) PRIMARY KEY
		)
	`)
	if err != nil {
		return err
	}

	log.Println("Database initialized, tables created.")
	return nil
}

// New returns an instance of Database
// Use this when you don't want to clear the DB
func New() *Database {
	db := dbConn(dbUser, dbPass, dbIP, dbPort, dbName)

	return &Database{
		db: db,
	}
}
