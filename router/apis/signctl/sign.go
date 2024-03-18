package signctl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sign/models/user"
	"sign/service/sign_service"
	"sign/tool/e"
)

func SignUp(c *gin.Context) {
	appG := e.Gin{C: c}
	in := &user.SignUpRequest{}
	if err := c.ShouldBindJSON(in); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	err := sign_service.NewSignService().SignUp(in)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func SignIn(c *gin.Context) {
	appG := e.Gin{C: c}
	in := &user.SignInRequest{}
	if err := c.ShouldBindJSON(in); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	info, err := sign_service.NewSignService().SignIn(in)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, info)
}
