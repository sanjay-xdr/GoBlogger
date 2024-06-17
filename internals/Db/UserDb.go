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
	stmt := `insert into users ("firstName" , "lastName","email","password",created_at , updated_at) values ($1,$2,$3,$4,$5,$6) returning id`
	_, err := m.DB.ExecContext(ctx, stmt, userData.FirstName, userData.LastName, userData.Email, userData.Password, userData.CreatedAt, userData.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgresDBRepo) GetUserById(id int) (models.User, error) {

	fmt.Println("Getting user by id")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT  "firstName", "lastName",email, created_at, updated_at FROM users WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, stmt, id)
	var user models.User
	err := row.Scan(&user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		fmt.Println(err, " Getting this error")
		return models.User{}, err
	}

	fmt.Println("Details of user is below")
	fmt.Print(user)
	return user, nil

}
