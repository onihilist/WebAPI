package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DatabaseHealthCheck(db *sql.DB) {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	actualTables := make(map[string]struct{})
	expectedTables := map[string]struct{}{
		"sqlite_sequence": {},
		"users":           {},
	}

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		actualTables[tableName] = struct{}{}
	}

	for expectedTable := range expectedTables {
		if _, exists := actualTables[expectedTable]; !exists {
			fmt.Printf("Missing expected table: %s\n", expectedTable)
		} else {
			fmt.Printf("Table exists: %s\n", expectedTable)
		}
	}

	for actualTable := range actualTables {
		if _, exists := expectedTables[actualTable]; !exists {
			fmt.Printf("Unexpected table found: %s\n", actualTable)
		}
	}
}

func DoRequest(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	res, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err) // Make an utils module for logs
		return nil
	}
	return res
}

func DoRequestRow(db *sql.DB, query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
}

func DatabaseDisconnect(db *sql.DB) {
	defer db.Close()
}
