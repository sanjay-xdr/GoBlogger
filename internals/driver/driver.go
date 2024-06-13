package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" //blank indentifier " _ " means we dont need the code directly in our code but need it indirectly
)

func ConnectDb(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {

		return nil, err
	}

	err = testDb(db)
	if err != nil {

		return nil, err
	}

	return db, nil

}

func testDb(db *sql.DB) error {

	err := db.Ping()

	if err != nil {
		return err
	}
	fmt.Println("Database ping successful")
	return nil

}
