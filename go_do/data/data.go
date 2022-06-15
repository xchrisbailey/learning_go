package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func DbConnect() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./go_do.db")
	if err != nil {
		log.Fatalf("could not establish db connection: %+v", err)
	}

	return db, nil
}

func InitDatabase(db *sql.DB) {
	createTodoTable := `CREATE TABLE IF NOT EXISTS todo (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "task" TEXT,
        "added_on" TEXT,
        "due_on" TEXT,
        "done" INTEGER
    );`

	stmt, err := db.Prepare(createTodoTable)
	if err != nil {
		log.Fatalf("could not create todo table: %+v", err)
	}

	stmt.Exec()
}
