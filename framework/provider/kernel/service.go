package kernel

import (
	"github.com/karry-11110/ginx/framework/gin"
	"net/http"
)

// 引擎服务
type GinxKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewGinxKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &GinxKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *GinxKernelService) HttpEngine() http.Handler {
	return s.engine
}
