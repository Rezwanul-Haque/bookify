package impl

import (
	"bookify/app/repository"
	"bookify/app/svc"
	"bookify/infra/serializers"
)

type system struct {
	repo repository.ISystem
}

func NewSystemService(sysrepo repository.ISystem) svc.ISystem {
	return &system{
		repo: sysrepo,
	}
}

func (sys *system) GetHealth() (*serializers.HealthResp, error) {
	resp := serializers.HealthResp{}

	// check db
	dbOnline, err := sys.repo.DBCheck()
	resp.DBOnline = dbOnline

	if err != nil {
		return &resp, err
	}

	return &resp, nil
}
