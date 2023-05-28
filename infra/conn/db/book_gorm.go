package db

import (
	"bookify/app/domain"
	"bookify/infra/conn/db/models"
	"bookify/infra/errors"
	"gorm.io/gorm"
)

func (dc DatabaseClient) Create(dbook *domain.Book) (*domain.Book, *errors.RestErr) {
	cc := models.Book{
		Title:           dbook.Title,
		Author:          dbook.Author,
		PublicationYear: dbook.PublicationYear,
	}

	res := dc.DB.Model(&models.Book{}).Create(&cc)

	if res.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	dbook.ID = cc.ID

	return dbook, nil
}

func (dc DatabaseClient) List() ([]*domain.Book, *errors.RestErr) {
	var books []*domain.Book
	err := dc.DB.Model(&models.Book{}).Find(&books).Error

	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	return books, nil
}

func (dc DatabaseClient) Get(bookID uint) (*domain.Book, *errors.RestErr) {
	var book *domain.Book
	err := dc.DB.Model(&models.Book{}).First(&book, bookID).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("resource not found")
	}
	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	return book, nil
}

func (dc DatabaseClient) Update(dbook *domain.Book) *errors.RestErr {
	cc := models.Book{
		ID:              dbook.ID,
		Title:           dbook.Title,
		Author:          dbook.Author,
		PublicationYear: dbook.PublicationYear,
	}

	res := dc.DB.Model(&models.Book{}).Where("id = ?", cc.ID).Updates(&cc)

	if res.Error != nil {
		return errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	return nil
}

func (dc DatabaseClient) Delete(bookID uint) *errors.RestErr {
	err := dc.DB.Delete(&models.Book{}, bookID).Error

	if err != nil {
		return errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	return nil
}
