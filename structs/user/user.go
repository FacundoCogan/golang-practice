package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() { // se convierte en un metodo del tipo recibido
	fmt.Printf("Usuario: %s %s | Nacimiento: %s | Creado: %s\n",
		u.firstName,
		u.lastName,
		u.birthDate,
		u.createdAt.Format("02/01/2006 15:04:05"),
	)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
