package services

import (
	"os"

	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/contract"
)

// GinxConsoleLog 代表控制台输出
type GinxConsoleLog struct {
	GinxLog
}

// NewGinxConsoleLog 实例化GinxConsoleLog
func NewGinxConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &GinxConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
