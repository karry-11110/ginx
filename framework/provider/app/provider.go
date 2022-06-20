package app

import (
    "github.com/karry-11110/ginx/framework"
    "github.com/karry-11110/ginx/framework/contract"
)

// GinxAppProvider 提供App的具体实现方法
type GinxAppProvider struct {
    BaseFolder string
}

// Register 注册GinxApp方法
func (h *GinxAppProvider) Register(container framework.Container) framework.NewInstance {
    return NewGinxApp
}

// Boot 启动调用
func (h *GinxAppProvider) Boot(container framework.Container) error {
    return nil
}

// IsDefer 是否延迟初始化
func (h *GinxAppProvider) IsDefer() bool {
    return false
}

// Params 获取初始化参数
func (h *GinxAppProvider) Params(container framework.Container) []interface{} {
    return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *GinxAppProvider) Name() string {
    return contract.AppKey
}
