package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poin4003/eCommerce_golang_api/internal/service"
	"github.com/poin4003/eCommerce_golang_api/internal/vo"
	"github.com/poin4003/eCommerce_golang_api/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid)
		return
	}

	fmt.Printf("Email param: %s", params.Email)

	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
