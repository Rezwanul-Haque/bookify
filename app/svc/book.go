package svc

import (
	"bookify/infra/errors"
	"bookify/infra/serializers"
)

type IBook interface {
	Create(book *serializers.BookReq) (*serializers.BookResp, *errors.RestErr)
	List() ([]*serializers.BookResp, *errors.RestErr)
	Get(id uint) (*serializers.BookResp, *errors.RestErr)
	Update(book *serializers.BookReq) *errors.RestErr
	Delete(id uint) *errors.RestErr
}
