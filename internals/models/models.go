package models

import "time"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Blog struct {
	Id         int
	Heading    string
	SubHeading string
	Content    string
	UserId     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Comment struct {
	Id        int
	Comment   string
	UserId    int
	BlogId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
