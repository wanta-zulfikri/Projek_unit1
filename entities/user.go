package entities

import "time"

type User struct {
	Id       int
	Username string
	Password string
	Role     string
}

type Log struct {
	OldUsername string
	NewUsername string
	Date        time.Time
}
