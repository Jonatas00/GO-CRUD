package controller

import "github.com/gin-gonic/gin"

type userController struct {
}

func NewUserController() userController {
	return userController{}
}

func (u *userController) GetUsers(ctx *gin.Context) {

}
