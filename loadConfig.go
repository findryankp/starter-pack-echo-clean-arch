package main

import (
	"cleanarc/apps/configs"
	"cleanarc/apps/database"
	"cleanarc/apps/routes"

	"github.com/labstack/echo/v4"
)

func loadConfigs() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	e := echo.New()
	routes.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
