package main

import (
	"github.com/Jonatas00/GO-anime-list/internal/config"
	"github.com/Jonatas00/GO-anime-list/internal/database"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	app := gin.Default()

	app.Routes()
	app.Run(":9090")
}
