package services

import (
	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/contract"
	"io"
)

type GinxCustomLog struct {
	GinxLog
}

func NewGinxCustomLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &GinxConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c
	return log, nil
}
