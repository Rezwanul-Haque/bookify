// Package classification bookify system API.
//
// the purpose of this service is to provides APIs related to book service
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /api
//	Version: v1.0.0
//	License: None
//	Contact: Rezwanul-Haque<rezwanul.cse@gmail.com>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package openapi

import (
	"bookify/infra/errors"
	"bookify/infra/serializers"
)

// Generic error message
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body errors.RestErr `json:"error_response"`
}

type genericSuccessResponse struct {
	// example: resource created
	Message string `json:"message"`
}

// returns a message
// swagger:response genericSuccessResponse
type genericSuccessResponseWrapper struct {
	// in: body
	genericSuccessResponse `json:"message"`
}

// Payload for create a book
// swagger:parameters Create
type bookPayloadWrapper struct {
	// in:body
	Body serializers.BookReq
}

// response after a book created
// swagger:response BookCreatedResponse
type bookCreateRespWrapper struct {
	// in:body
	Body serializers.BookResp
}

// List all the books of a company
// swagger:response BookResponse
type booksRespWrapper struct {
	// in:body
	Body []*serializers.BookResp
}
