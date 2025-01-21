package databases

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "../../database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
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
