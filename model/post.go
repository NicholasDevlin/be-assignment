package model

import "time"

type User struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
}
