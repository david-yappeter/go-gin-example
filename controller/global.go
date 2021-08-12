package controller

import (
	"fmt"
	"myapp/service"
	"myapp/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	user service.UserService
}

var validate *validator.Validate

var (
	userService service.UserService = service.NewUserService()
	ctrl        *controller         = newController(userService)
)

var (
	gin_INTERNAL_SYSTEM_ERROR = ginWrapper(nil, fmt.Errorf("internal system error"))
	gin_BAD_REQUEST           = ginWrapper(nil, fmt.Errorf("bad request"))
)

func newController(userService service.UserService) *controller {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		user: userService,
	}
}

func ginWrapper(data interface{}, errParam ...error) gin.H {
	var errString *string = nil
	if len(errParam) > 0 {
		temp := errParam[0].Error()
		errString = &temp
	}

	return gin.H{
		"data":      data,
		"error_msg": errString,
	}
}
