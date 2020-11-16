package db

import (
	"log"
	"database/sql"
	"fmt"
)

func dbConn(dbUser string, dbPass string, dbIP string, dbPort int, dbName string) *sql.DB {
	dbDriver := "mysql"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&autocommit=true", dbUser, dbPass, dbIP, dbPort, dbName)

	db, err := sql.Open(dbDriver, dbURI)

	if err != nil {
		log.Println(err)
		log.Fatal("Error connecting to database.")
	}

	log.Println("Database connection successful!")

	return db
}

// PingDB allows you to ping the database
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