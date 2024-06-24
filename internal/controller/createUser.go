package controller

import (
	"fmt"
	"log"

	"github.com/Jonatas00/GO-CRUD/internal/config/validation"
	"github.com/Jonatas00/GO-CRUD/internal/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var UserRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("There are some incorret fields, error=%s \n", err.Error())
		errRest := validation.ValidateUserError(err)

		ctx.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(UserRequest)

	ctx.JSON(200, UserRequest)
}
