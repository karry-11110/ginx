package tests

import (
	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/provider/app"
	"github.com/karry-11110/ginx/framework/provider/env"
)

const (
	BasePath = "/Users/wangkun/Documents/UGit/coredemo/"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewGinxContainer()
	// 绑定App服务提供者
	container.Bind(&app.GinxAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.GinxTestingEnvProvider{})
	return container
}
