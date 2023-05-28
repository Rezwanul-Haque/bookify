package domain

import (
	"bookify/infra/errors"
	"time"
)

type Book struct {
	ID              uint
	Title           string
	Author          string
	PublicationYear string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type IBook interface {
	Create(book *Book) (*Book, *errors.RestErr)
	List() ([]*Book, *errors.RestErr)
	Get(id uint) (*Book, *errors.RestErr)
	Update(book *Book) *errors.RestErr
	Delete(id uint) *errors.RestErr
}
