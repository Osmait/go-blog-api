package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepositoy interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
}

func NewUser(email string, password string) (*User, error) {
	idU, err := uuid.NewUUID()
	if err != nil {
		return &User{}, err
	}

	return &User{
		Id:       idU,
		Email:    email,
		Password: password,
	}, nil

}

func (u User) ID() uuid.UUID {
	return u.Id

}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) SetID(id uuid.UUID) {
	u.Id = id

}

func (u User) SetEmail(email string) {
	u.Email = email
}

func (u User) SetPassword(password string) {
	u.Password = password
}
