package Db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/sanjay-xdr/goblogger/internals/models"
)

func (m *PostgresDBRepo) InsertIntoBlogs(blog models.Blog) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into blogs ("heading" , "subHeading","content","userId",created_at , updated_at) values ($1,$2,$3,$4,$5,$6) returning id`
	_, err := m.DB.ExecContext(ctx, stmt, blog.Heading, blog.SubHeading, blog.Content, blog.UserId, blog.CreatedAt, blog.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgresDBRepo) UpdateBlogs(blog models.Blog, id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `update blogs set heading=$1 , "subHeading"=$2 , content=$3 , "userId"=$4 , created_at=$5 , updated_at=$6 where id=$7`
	_, err := m.DB.ExecContext(ctx, stmt, blog.Heading, blog.SubHeading, blog.Content, blog.UserId, blog.CreatedAt, blog.UpdatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgresDBRepo) GetAllBlogs() ([]models.Blog, error) {

	var bloglist []models.Blog

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select id, heading, "subHeading" , content , "userId" , created_at , updated_at from blogs`
	row, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		log.Fatal("Not able to read rows")
	}

	for row.Next() {
		var blog models.Blog
		err := row.Scan(&blog.Id, &blog.Heading, &blog.SubHeading, &blog.Content, &blog.UserId, &blog.CreatedAt, &blog.UpdatedAt)
		if err != nil {
			log.Fatal("Not able to read rows")
		}
		bloglist = append(bloglist, blog)
	}

	fmt.Print(bloglist)
	return bloglist, nil

}
func (m *PostgresDBRepo) GetBlogById(id int) (models.Blog, error) {

	var blog models.Blog
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select id, heading, "subHeading" , content , "userId" , created_at , updated_at from blogs where id=$1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	err := row.Scan(&blog.Heading, &blog.SubHeading, &blog.Content, &blog.UserId, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		if sql.ErrNoRows == err {
			log.Fatal("THere is no row")
		}
		log.Fatal("Something went wrong in dbfun", err)
	}
	return blog, nil

}

func (m *PostgresDBRepo) DeleteBlogById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `DELETE FROM blogs WHERE id = $1`
	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
