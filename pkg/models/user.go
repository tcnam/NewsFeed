package models

import "time"

type User struct {
	UserID       int       `json:"id"`
	Name         string    `json:"name"`
	dob          time.Time `json:"date_of_birth"`
	phone_number string    `json:"phone_number"`
	email        string    `json:"json"`
	password     string    `json:"password"`
	active       bool      `json:"active"`
	created_at   time.Time `json:"created_at"`
}

type UserService interface {
	Validate(user *User) error
	ValidateAge(user *User) bool
	CreateUser(user *User) (*User, error)
	Findall() ([]User, error)
}

type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error
}

func CreateUser(user *User) (*User, error) {

}
