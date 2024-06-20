package main

import (
	"fmt"

	"github.com/Jonatas00/GO-CRUD/internal/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	app := gin.Default()

	app.Run(fmt.Sprintf(":%s", config.AppCfg.Port))
}
