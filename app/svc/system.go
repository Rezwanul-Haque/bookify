package svc

import "bookify/infra/serializers"

type ISystem interface {
	GetHealth() (*serializers.HealthResp, error)
}
