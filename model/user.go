package model

import "time"

//User user
type User struct {
	ID             int64
	UserName       string
	Email          string
	Phone          string
	EmailValidated int
	PhoneValidated int
	CreateTime     time.Time
	LastActiveTime time.Time
	Status         int
	Profile        string
}
