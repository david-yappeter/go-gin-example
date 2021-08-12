package controller

import (
	"fmt"
	"myapp/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserGetByUUID(c *gin.Context) {
	uuID := c.Params.ByName("uuID")
	user, err := ctrl.user.GetByID(uuID)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin_INTERNAL_SYSTEM_ERROR)
	} else {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, ginWrapper(user, err))
		} else {
			c.JSON(http.StatusOK, ginWrapper(user))
		}
	}
}

func UserGetAll(c *gin.Context) {
	user, err := ctrl.user.GetAll()
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin_INTERNAL_SYSTEM_ERROR)
	} else {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, ginWrapper(user, err))
		} else {
			c.JSON(http.StatusOK, ginWrapper(user))
		}
	}
}

func UserCreate(c *gin.Context) {
	var userForm entity.NewUserForm
	if err := c.ShouldBindJSON(&userForm); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin_BAD_REQUEST)
	} else {
		user, err := ctrl.user.Create(entity.User(userForm))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin_INTERNAL_SYSTEM_ERROR)
		} else {
			c.JSON(http.StatusOK, ginWrapper(user))
		}
	}
}

// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-urlencoded
// @Produce json
// @Param UserData body entity.UserLoginForm true "Email & Password"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.JWT
// @Router /login [post]
func UserLogin(c *gin.Context) {
	var userLoginForm entity.UserLoginForm
	if err := c.ShouldBindJSON(&userLoginForm); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin_BAD_REQUEST)
	} else {
		token, err := ctrl.user.Login(userLoginForm)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, ginWrapper(nil, fmt.Errorf("email not found")))
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusOK, ginWrapper(nil, fmt.Errorf("wrong password")))
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin_INTERNAL_SYSTEM_ERROR)
		} else {
			c.JSON(http.StatusOK, ginWrapper(gin.H{"type": "Bearer",
				"token": token}))
		}
	}
}
