package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/onihilist/WebAPI/pkg/utils"
)

func DatabaseConnect() *sql.DB {
	dsn := "appuser:letmein@tcp(mariadb:3306)/appdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		utils.LogFatal("[MariaDB] - %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		utils.LogFatal("[MariaDB] - %s", err.Error())
	}

	return db
}

func DatabaseHealthCheck(db *sql.DB) {
	rows, err := db.Query("SELECT table_name FROM maria_schema WHERE table_type='BASE TABLE'")
	if err != nil {
		utils.LogFatal("[MariaDB] - %s", err.Error())
	}
	defer rows.Close()

	actualTables := make(map[string]struct{})
	expectedTables := map[string]struct{}{
		"maria_schema": {},
		"users":        {},
		"permissions":  {},
	}

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			utils.LogFatal("[MariaDB] - %s", err.Error())
		}
		actualTables[tableName] = struct{}{}
	}

	for expectedTable := range expectedTables {
		if _, exists := actualTables[expectedTable]; !exists {
			utils.LogError("[MariaDB] - Missing expected table: %s\n", expectedTable)
		} else {
			utils.LogSuccess("[MariaDB] - Table exists: %s\n", expectedTable)
		}
	}

	for actualTable := range actualTables {
		if _, exists := expectedTables[actualTable]; !exists {
			utils.LogWarning("[MariaDB] - Unexpected table found: ", actualTable)
		}
	}
}

func DoRequest(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	res, err := db.Query(query, args...)
	if err != nil {
		utils.LogError("[MariaDB] - %s", err.Error())
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
