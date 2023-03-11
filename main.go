package main

import (
	"immersiveApp/app/configs"
	"immersiveApp/app/database"
	"immersiveApp/app/router"
	"immersiveApp/docs"

	"github.com/labstack/echo/v4"
)

func main() {
	docs.InitSwagger()
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	e := echo.New()
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
