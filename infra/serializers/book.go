package serializers

import "time"

type BookReq struct {
	ID              uint `json:"-"`
	Title           string
	Author          string
	PublicationYear string
}

type BookResp struct {
	ID              uint
	Title           string
	Author          string
	PublicationYear string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
