package main

import (
	"fmt"
	"price-check/app/models"
	"price-check/config"
	"price-check/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}
