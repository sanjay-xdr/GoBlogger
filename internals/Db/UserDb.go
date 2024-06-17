package Db

import (
	"context"
	"fmt"
	"time"

	"github.com/sanjay-xdr/goblogger/internals/models"
)

func (m *PostgresDBRepo) InsertIntoUser(userData models.User) error {

	fmt.Print("adding into DB")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into tester ("firstName" , "lastName","email","password",created_at , updated_at) values ($1,$2,$3,$4,$5,$6) returning id`
	_, err := m.DB.ExecContext(ctx, stmt, userData.FirstName, userData.LastName, userData.Email, userData.Password, userData.CreatedAt, userData.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
