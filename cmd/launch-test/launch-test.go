package main

import (
	"fmt"

	"gitlab.utc.fr/mennynat/ia04-project/serveur"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

func main() {

	leLogger := logger.ConsoleLogger{}

	loggers := logger.Loggers{}
	loggers.AjouterLogger(&leLogger)

	server := serveur.NewRestServerAgent("localhost:8080")
	server.Start()
	fmt.Scanln()
}
