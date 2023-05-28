package controllers

import (
	"bookify/app/svc"
	"bookify/app/utils/msgutil"
	"bookify/infra/errors"
	"bookify/infra/logger"
	"bookify/infra/serializers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type book struct {
	lc  logger.LogClient
	svc svc.IBook
}

// NewBookController will initialize the controllers
func NewBookController(grp interface{}, svc svc.IBook) {
	uc := &book{
		svc: svc,
	}

	g := grp.(*echo.Group)

	g.POST("/v1/book", uc.Create)
	g.GET("/v1/book", uc.List)
	g.GET("/v1/book/:id", uc.Get)
	g.PUT("/v1/book/:id", uc.Update)
	g.DELETE("/v1/book/:id", uc.Delete)
}

// swagger:route POST /v1/book Book Create
// Create a new book
// responses:
//	201: BookCreatedResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
//

// Create handles POST requests and create a new book
func (ctrlr *book) Create(c echo.Context) error {
	var bookReq *serializers.BookReq

	if err := c.Bind(&bookReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	result, saveErr := ctrlr.svc.Create(bookReq)
	if saveErr != nil {
		return c.JSON(saveErr.Status, saveErr)
	}

	return c.JSON(http.StatusCreated, result)
}

// swagger:route POST /v1/book Book List
// Fetch all the books
// responses:
//	200: BookResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
//

// List handles GET requests and fetch all the books
func (ctrlr *book) List(c echo.Context) error {

	resp, getErr := ctrlr.svc.List()

	if getErr != nil {
		return c.JSON(getErr.Status, getErr)
	}

	return c.JSON(http.StatusOK, &resp)
}

// swagger:route POST /v1/book Book Get
// Fetch a book
// responses:
//	200: BookCreatedResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
//

// Get handles GET requests and fetch a book
func (ctrlr *book) Get(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, errors.NewBadRequestError("id can't be empty"))
	}

	resp, getErr := ctrlr.svc.Get(uint(bookID))
	if getErr != nil {
		return c.JSON(getErr.Status, getErr)
	}

	return c.JSON(http.StatusOK, &resp)
}

// swagger:route PUT /v1/book Book Update
// Update a book
// responses:
//	200: genericSuccessResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
//

// Update handles Put requests and update a book
func (ctrlr *book) Update(c echo.Context) error {
	var bookReq *serializers.BookReq
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, errors.NewBadRequestError("id can't be empty"))
	}
	if err := c.Bind(&bookReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	bookReq.ID = uint(bookID)

	updateErr := ctrlr.svc.Update(bookReq)
	if updateErr != nil {
		return c.JSON(updateErr.Status, updateErr)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": msgutil.EntityUpdateSuccessMsg("book")})
}

// swagger:route DELETE /v1/book Book Delete
// Delete a book
// responses:
//	204:
//	400: errorResponse
//	500: errorResponse
//

// Delete handles delete requests and delete a book
func (ctrlr *book) Delete(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, errors.NewBadRequestError("id can't be empty"))
	}
	delErr := ctrlr.svc.Delete(uint(bookID))

	if delErr != nil {
		return c.JSON(delErr.Status, delErr)
	}

	return c.NoContent(http.StatusNoContent)
}
