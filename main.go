package main

import (
	"log"
	"price-check/app/controllers"
	"price-check/config"
	"price-check/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
