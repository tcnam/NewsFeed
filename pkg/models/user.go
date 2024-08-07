package models

import "time"

type User struct {
	UserID       int
	Name         string
	dob          time.Time
	phone_number string
	email        string
	password     string
	active       bool
	created_at   time.Time
}

type UserService interface {
	Validate(user *User) error
	ValidateAge(user *User) bool
	Create(user *User) (*User, error)
	Findall() ([]User, error)
}

type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error
}
