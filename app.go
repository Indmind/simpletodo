package main

import (
	"database/sql"

	"github.com/indmind/simpletodo/handlers"

	_ "github.com/mattn/go-sqlite3"
	echo "gopkg.in/echo.v3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.Static("/", "public")

	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Start(":8080")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("DB nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
