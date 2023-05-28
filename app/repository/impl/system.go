package impl

import (
	"bookify/app/repository"
	"bookify/infra/conn/db"
	"context"
)

type system struct {
	ctx context.Context
	DB  db.DatabaseClient
}

// NewSystemRepository will create an object that represent the System.Repository implementations
func NewSystemRepository(ctx context.Context, dbc db.DatabaseClient) repository.ISystem {
	return &system{
		ctx: ctx,
		DB:  dbc,
	}
}

func (r *system) DBCheck() (bool, error) {
	dB, _ := r.DB.DB.DB()
	if err := dB.Ping(); err != nil {
		return false, err
	}

	return true, nil
}
