package controllers

import (
	"bookify/app/svc"
	"bookify/infra/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type system struct {
	svc svc.ISystem
}

// NewSystemController will initialize the controllers
func NewSystemController(grp interface{}, sysSvc svc.ISystem) {
	pc := &system{
		svc: sysSvc,
	}

	g := grp.(*echo.Group)

	g.GET("/v1", pc.Root)
	g.GET("/v1/h34l7h", pc.Health)
}

// Root will let you see what you can slash 🐲
func (ctrlr *system) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "bookify architecture backend! let's play!!"})
}

// swagger:route GET /v1/h34l7h Health will let you know the heart beats ❤️
// Return a message
// responses:
//	200: genericSuccessResponse

// Health will let you know the heart beats ❤️
func (ctrlr *system) Health(c echo.Context) error {
	resp, err := ctrlr.svc.GetHealth()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ErrSomethingWentWrong)
	}
	return c.JSON(http.StatusOK, resp)
}
