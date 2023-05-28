package impl

import (
	"bookify/app/domain"
	"bookify/app/repository"
	"bookify/app/svc"
	"bookify/app/utils/methodutil"
	"bookify/infra/errors"
	"bookify/infra/logger"
	"bookify/infra/serializers"
	"context"
)

type book struct {
	ctx   context.Context
	lc    logger.LogClient
	brepo repository.IBook
}

func NewUsersService(ctx context.Context, lc logger.LogClient, brepo repository.IBook) svc.IBook {
	return &book{
		ctx:   ctx,
		lc:    lc,
		brepo: brepo,
	}
}

func (b *book) Create(req *serializers.BookReq) (*serializers.BookResp, *errors.RestErr) {
	var dbook *domain.Book
	_ = methodutil.StructToStruct(req, &dbook)

	result, saveErr := b.brepo.Create(dbook)
	if saveErr != nil {
		return nil, saveErr
	}

	var resp *serializers.BookResp
	_ = methodutil.StructToStruct(result, &resp)

	return resp, nil
}

func (b *book) List() ([]*serializers.BookResp, *errors.RestErr) {
	result, saveErr := b.brepo.List()
	if saveErr != nil {
		return nil, saveErr
	}

	var resp []*serializers.BookResp
	_ = methodutil.StructToStruct(result, &resp)

	return resp, nil
}

func (b *book) Get(bookID uint) (*serializers.BookResp, *errors.RestErr) {
	result, getErr := b.brepo.Get(bookID)
	if getErr != nil {
		return nil, getErr
	}

	var resp *serializers.BookResp
	_ = methodutil.StructToStruct(result, &resp)

	return resp, nil
}

func (b *book) Update(req *serializers.BookReq) *errors.RestErr {
	var updateBook *domain.Book
	_ = methodutil.StructToStruct(req, &updateBook)

	updateBook.ID = req.ID

	updErr := b.brepo.Update(updateBook)
	if updErr != nil {
		return updErr
	}
	return nil
}

func (b *book) Delete(bookID uint) *errors.RestErr {
	delErr := b.brepo.Delete(bookID)
	if delErr != nil {
		return delErr
	}
	return nil
}
