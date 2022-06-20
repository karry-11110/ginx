# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/karry-11110/ginx/framework/gin"
	"github.com/karry-11110/ginx/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
