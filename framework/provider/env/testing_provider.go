package env

import (
	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/contract"
)

type GinxTestingEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *GinxTestingEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewGinxTestingEnv
}

// Boot will called when the service instantiate
func (provider *GinxTestingEnvProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GinxTestingEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *GinxTestingEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *GinxTestingEnvProvider) Name() string {
	return contract.EnvKey
}
