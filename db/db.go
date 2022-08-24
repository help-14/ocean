package db

import (
	"database/sql"
	"log"
	"os"
)

const dbPath = "./ocean.db"
const serviceTag = "database"

var db *sql.DB = nil

func Connect() error {
	// Create SQLite file
	dbExist := true
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		dbExist = false
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		file.Close()
	}

	// Open the created SQLite File
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// Defer Closing the database
	defer db.Close()

	// Create Database Tables
	if !dbExist {
		err = createTable()
		if err != nil {
			return err
		}
	}
	return nil
}

func checkDb() {
	if db == nil {
		log.Fatalln("Database is not setup yet!")
	}
}

func createTable() error {
	checkDb()
	createTableSQL := []string{
		`CREATE TABLE logs (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
			"job" VARCHAR NOT NULL,
			"path" VARCHAR NOT NULL,
			"time" DATETIME DEFAULT (datetime('now','localtime')),
			"success" BOOLEAN,
			"error" TEXT);`,
		`CREATE TABLE syslogs (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
			"tag" VARCHAR NOT NULL,
			"time" DATETIME DEFAULT (datetime('now','localtime')),
			"error" TEXT);`,
	}
	// `CREATE INDEX idx_contacts_name ON logs (first_name, last_name);`

	for _, sql := range createTableSQL {
		statement, err := db.Prepare(sql)
		if err != nil {
			log.Fatalln("Prepare table failed: ", err.Error())
			return err
		}

		_, err = statement.Exec()
		if err != nil {
			log.Fatalln("Create table failed: ", err.Error())
			return err
		}
	}

	return nil
}
