package Db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/sanjay-xdr/goblogger/internals/models"
)

func (m *PostgresDBRepo) InsertIntoComments(comment models.Comment) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into comments ("comment" , "userId","blogid",created_at , updated_at) values ($1,$2,$3,$4,$5) returning id`
	_, err := m.DB.ExecContext(ctx, stmt, comment.Comment, comment.UserId, comment.BlogId, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgresDBRepo) UpdateComments(comment models.Comment, id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `update comments set comment=$1 , "userId"=$2 , "blogid"=$3 ,created_at=$4 , updated_at=$5 where id=$6`
	_, err := m.DB.ExecContext(ctx, stmt, comment.Comment, comment.UserId, comment.BlogId, comment.CreatedAt, comment.UpdatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgresDBRepo) GetCommentByBlogId(id int) (models.Comment, error) {

	var cmt models.Comment
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select  comment, "userId" ,"blogid" ,created_at , updated_at from comments where blogid=$1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	err := row.Scan(&cmt.Comment, &cmt.UserId, &cmt.BlogId, &cmt.CreatedAt, &cmt.UpdatedAt)
	if err != nil {
		if sql.ErrNoRows == err {
			log.Fatal("THere is no row")
		}
		log.Fatal("Something went wrong", err)
	}
	return cmt, nil

}

func (m *PostgresDBRepo) DeleteCommentById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `DELETE FROM comments WHERE id = $1`
	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
