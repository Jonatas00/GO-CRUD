package main

import (
	"fmt"

	"github.com/Jonatas00/GO-CRUD/internal/config"
	"github.com/Jonatas00/GO-CRUD/internal/database"
	"github.com/Jonatas00/GO-CRUD/internal/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()

	database.Connect()
	database.GenerateMigrations()
}

func main() {
	app := gin.Default()

	routes.RunRoutes(&app.RouterGroup)

	app.Run(fmt.Sprintf(":%s", config.AppCfg.APP_PORT))
}
