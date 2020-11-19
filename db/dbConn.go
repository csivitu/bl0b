package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func dbConn(dbUser string, dbPass string, dbIP string, dbPort int, dbName string) *sqlx.DB {
	dbDriver := "mysql"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&autocommit=true&parseTime=true", dbUser, dbPass, dbIP, dbPort, dbName)

	db, err := sqlx.Open(dbDriver, dbURI)

	if err != nil {
		log.Println(err)
		log.Fatal("Error connecting to database.")
	}

	return db
}

// Ping allows you to ping the database
// and check if connection is possible
func (DB *Database) Ping() {
	err := DB.db.Ping()

	if err != nil {
		log.Fatal("Could not ping DB!")
	}
}

// Close closes the connetion to the database,
// can call this with defer in main
func (DB *Database) Close() {
	DB.db.Close()
}
