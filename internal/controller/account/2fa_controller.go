package account

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poin4003/eCommerce_golang_api/internal/model"
	"github.com/poin4003/eCommerce_golang_api/internal/service"
	"github.com/poin4003/eCommerce_golang_api/internal/utils/context"
	"github.com/poin4003/eCommerce_golang_api/pkg/response"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

// User setup Two Factor Authentication documentation
// @Summary setup Two Factor Authentication godoc
// @Description When user want setup Two Factor Authentication
// @Tags account 2fa
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param payload body model.SetupTwoFactorAuthInput true "payload"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /user/two_factor/setup [post]
func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid setupTwoFactorAuth")
		return
	}

	// get userId from uuid (token)
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not valid")
		return
	}
	log.Println("UserId: ", userId)
	params.UserId = int32(userId)

	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeResult, nil)
}
