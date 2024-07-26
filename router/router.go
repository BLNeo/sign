package router

import (
	"github.com/gin-gonic/gin"
	"sign/middleware"
	"sign/router/apis/signctl"
	"sign/router/apis/testctl"
)

// InitRouter initialize routing information
func InitRouter(r *gin.Engine) {
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CORSMiddleware())
	testRoute(r.Group("/test"))
	signRoute(r.Group(""))
}

func testRoute(rg *gin.RouterGroup) {
	rg.GET("", testctl.GetTest)
}

func signRoute(rg *gin.RouterGroup) {
	rg.POST("/sign_in", signctl.SignIn)
	rg.POST("/sign_up", signctl.SignUp)
}
