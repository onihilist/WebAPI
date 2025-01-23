package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() *sql.DB {
	dsn := "appuser:letmein@tcp(mariadb:3306)/appdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func DatabaseHealthCheck(db *sql.DB) {
	rows, err := db.Query("SELECT table_name FROM maria_schema WHERE table_type='BASE TABLE'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	actualTables := make(map[string]struct{})
	expectedTables := map[string]struct{}{
		"maria_schema": {},
		"users":        {},
	}

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		actualTables[tableName] = struct{}{}
	}

	fmt.Printf("Actual tables : %s\n", actualTables)

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
