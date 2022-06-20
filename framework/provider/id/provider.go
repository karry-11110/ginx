package id

import (
	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/contract"
)

type GinxIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *GinxIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewGinxIDService
}

// Boot will called when the service instantiate
func (provider *GinxIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GinxIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *GinxIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *GinxIDProvider) Name() string {
	return contract.IDKey
}
