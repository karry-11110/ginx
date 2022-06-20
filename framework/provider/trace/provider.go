package trace

import (
	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/contract"
)

type GinxTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *GinxTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewGinxTraceService
}

// Boot will called when the service instantiate
func (provider *GinxTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GinxTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *GinxTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *GinxTraceProvider) Name() string {
	return contract.TraceKey
}
