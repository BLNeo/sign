package signctl

import (
	"github.com/gin-gonic/gin"
	"sign/models/user"
	"sign/service/sign_service"
	"sign/tool/response"
	"sign/tool/util"
)

func SignUp(c *gin.Context) {

	in := &user.SignUpRequest{}
	if err := util.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := sign_service.NewSignService().SignUp(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func SignIn(c *gin.Context) {
	in := &user.SignInRequest{}
	if err := util.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	info, err := sign_service.NewSignService().SignIn(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, info)
}
