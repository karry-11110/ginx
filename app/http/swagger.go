// Package http API.
// @title ginx
// @version 1.1
// @description ginx测试
// @termsOfService https://github.com/swaggo/swag

// @contact.name wangkun1
// @contact.email wangkun

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

import (
	_ "github.com/karry-11110/ginx/app/http/swagger"
)
