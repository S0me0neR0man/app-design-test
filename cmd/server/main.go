package main

import (
	"log"

	"applicationDesignTest/api"
	"applicationDesignTest/internals/infrastructure/storage/simple"
	"applicationDesignTest/internals/usecases/orders"
)

const (
	MaxRateLimit = 100
)

func main() {
	logger := log.Default()
	database := simple.NewOnSliceDatabase()
	orderUseCases := orders.NewOrderUseCases(database, database, logger)

	webApp := api.NewWebApp(orderUseCases, logger)

	webApp.Run()
}
