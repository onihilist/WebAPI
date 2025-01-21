package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("sqlite", "../../database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DatabaseHealthCheck(db *sql.DB) {
	tables, err := db.Query("SELECT name FROM sqlite_sequence WHERE type='table';")
	if err != nil {
		fmt.Println("Error (DatabaseHealthCheck) :", err)
	} else {
		fmt.Println(tables)
	}
}

func DoRequest(db *sql.DB, query string) *sql.Rows {
	res, err := db.Query(query)
	if err != nil {
		fmt.Println(err) // Make an utils module for logs
	}
	return res
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
