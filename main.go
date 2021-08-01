package main

import (
	"price-check/app/controllers"
	"price-check/config"
	"price-check/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
