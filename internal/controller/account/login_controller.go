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

// User login documentation
// @Summary User godoc
// @Description When user want to login
// @Tags account management
// @Accept json
// @Produce json
// @Param payload body model.LoginInput true "payload"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /user/login/ [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	var params model.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeRs, dataRs, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}

// User Registraion documentation
// @Summary User Registration
// @Description When user is registered send otp to email
// @Tags account management
// @Accept json
// @Produce json
// @Param payload body model.RegisterInput true "payload"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /user/register/ [post]
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

// VerifyOTP documentation
// @Summary Verify OTP login by User
// @Description Verify OTP login by User
// @Tags account management
// @Accept json
// @Produce json
// @Param payload body model.VerifyInput true "payload"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /user/verify_account/ [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidOTP, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}

// UpdatePasswordRegister documentation
// @Summary Update password register
// @Description Update password register
// @Tags account management
// @Accept json
// @Produce json
// @Param payload body model.UpdatePasswordRegisterInput true "payload"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /user/update_pass_register/ [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
