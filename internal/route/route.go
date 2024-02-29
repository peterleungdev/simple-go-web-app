package route

import "github.com/gin-gonic/gin"

func StartServer() {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) { ctx.String(200, "HELLO WORLD running go server") })
	route.GET("/ping", pong)

	route.Run(":8088")
}

func pong(ctx *gin.Context) {
	ctx.String(200, "pong")
}
