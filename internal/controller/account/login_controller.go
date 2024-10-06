package account

import (
	"github.com/gin-gonic/gin"
	"github.com/poin4003/eCommerce_golang_api/global"
	"github.com/poin4003/eCommerce_golang_api/internal/model"
	"github.com/poin4003/eCommerce_golang_api/internal/service"
	"github.com/poin4003/eCommerce_golang_api/pkg/response"
	"go.uber.org/zap"
)

type cUserLogin struct {
}

// management container Login User
var Login = new(cUserLogin)

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}

func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, nil)
}
