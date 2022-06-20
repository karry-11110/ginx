package id

import (
	"github.com/rs/xid"
)

type GinxIDService struct {
}

func NewGinxIDService(params ...interface{}) (interface{}, error) {
	return &GinxIDService{}, nil
}

func (s *GinxIDService) NewID() string {
	return xid.New().String()
}
