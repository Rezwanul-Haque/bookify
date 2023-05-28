package impl

import (
	"bookify/app/domain"
	"bookify/app/repository"
	"bookify/infra/conn/db"
	"bookify/infra/errors"
	"context"
)

type book struct {
	ctx context.Context
	DB  db.DatabaseClient
}

// NewBookRepository will create an object that represent the Book.Repository implementations
func NewBookRepository(ctx context.Context, dbc db.DatabaseClient) repository.IBook {
	return &book{
		ctx: ctx,
		DB:  dbc,
	}
}

func (b *book) Create(book *domain.Book) (*domain.Book, *errors.RestErr) {
	resp, saveErr := b.DB.Create(book)
	if saveErr != nil {
		return nil, saveErr
	}
	return resp, nil
}

func (b *book) List() ([]*domain.Book, *errors.RestErr) {
	resp, saveErr := b.DB.List()
	if saveErr != nil {
		return nil, saveErr
	}
	return resp, nil
}

func (b *book) Get(bookID uint) (*domain.Book, *errors.RestErr) {
	resp, getErr := b.DB.Get(bookID)
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}

func (b *book) Update(book *domain.Book) *errors.RestErr {
	updErr := b.DB.Update(book)
	if updErr != nil {
		return updErr
	}
	return nil
}

func (b *book) Delete(bookID uint) *errors.RestErr {
	delErr := b.DB.Delete(bookID)
	if delErr != nil {
		return delErr
	}
	return nil
}
