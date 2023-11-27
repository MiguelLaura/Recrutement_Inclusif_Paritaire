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
	//loggers.Log("cc", "comment", "ça", "va", "?")
	//leLogger.Log("cc", "comment", "ça", "va", "?")

	server := serveur.NewRestServerAgent(":8080")
	server.Start()
	fmt.Scanln()
}
