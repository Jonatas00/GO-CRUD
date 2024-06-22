package main

import (
	"fmt"
	"log"

	"github.com/Jonatas00/GO-CRUD/internal/config"
	"github.com/Jonatas00/GO-CRUD/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	app := gin.Default()

	routes.InitRoutes(&app.RouterGroup)

	if err := app.Run(fmt.Sprintf(":%s", config.AppCfg.Port)); err != nil {
		log.Fatal(err)
	}
}
