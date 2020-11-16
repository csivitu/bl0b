package db

import (
	"log"
	"database/sql"
	"fmt"
)

func dbConn(dbUser string, dbPass string, dbIP string, dbPort int, dbName string) *sql.DB {
	dbDriver := "mysql"
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbIP, dbPort, dbName)

	db, err := sql.Open(dbDriver, connString)

	if err != nil {
		log.Fatal("Error connecting to database.")
	}

	return db
}

func pingDB(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal("Could not ping DB!")
	}
}
