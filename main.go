package main

import (
	"github.com/fgkinus/fileManager/src"
	"github.com/fgkinus/fileManager/src/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	// read the env
	arguments := config.GetCMDArgs()
	config.ReadCompileTimeEnv()
	config.ReadRuntimeConfig(arguments.ConfigFilePath)
	// database connection
	configuration := config.GetConfig()
	err := config.ConnectToDb(configuration.Database.Uri)
	if err != nil {
		config.Logger.Fatal(err)
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  ${method}  ${uri}  ${status} \n",
	}))
	e.Use(middleware.Recover())

	// configure the app routes
	src.RoutesManager(e)
	// Start server
	err := e.Start(":1323")
	if err != nil {
		config.Logger.Fatal(err)
	}
}
