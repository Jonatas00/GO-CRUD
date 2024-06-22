package controller

import (
	"fmt"

	"github.com/Jonatas00/GO-CRUD/internal/config/restErr"
	"github.com/Jonatas00/GO-CRUD/internal/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var UserRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&UserRequest); err != nil {
		erro := restErr.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s \n", err.Error()),
		)

		ctx.JSON(erro.Code, erro)
		return
	}

	fmt.Println(UserRequest)

	ctx.JSON(200, UserRequest)
}
