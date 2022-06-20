package env

import (
    "github.com/karry-11110/ginx/framework"
    "github.com/karry-11110/ginx/framework/contract"
)

type GinxEnvProvider struct {
    Folder string
}

// Register registe a new function for make a service instance
func (provider *GinxEnvProvider) Register(c framework.Container) framework.NewInstance {
    return NewGinxEnv
}

// Boot will called when the service instantiate
func (provider *GinxEnvProvider) Boot(c framework.Container) error {
    app := c.MustMake(contract.AppKey).(contract.App)
    provider.Folder = app.BaseFolder()
    return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GinxEnvProvider) IsDefer() bool {
    return false
}

// Params define the necessary params for NewInstance
func (provider *GinxEnvProvider) Params(c framework.Container) []interface{} {
    return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *GinxEnvProvider) Name() string {
    return contract.EnvKey
}
