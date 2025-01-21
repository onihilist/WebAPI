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
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'") // Add an array to see if all the tables are created
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Table:", tableName)
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

func CreateUser() {
	// INSERT A USER INTO THE DATABASE
	// _, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "john@example.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func DumpAllUsers() {
	// DUMP ALL USERS AND PRINT THEM
	// rows, err := db.Query("SELECT id, name, email FROM users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var name, email string
	// 	err = rows.Scan(&id, &name, &email)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf(":User  %d, Name: %s, Email: %s\n", id, name, email)
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
