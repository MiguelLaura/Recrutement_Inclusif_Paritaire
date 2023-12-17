package main

import (
	"log"

	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

func main() {

	leLogger := logger.ConsoleLogger{}
	leLoggerF, err := logger.NewFileLogger("")

	if err != nil {
		log.Fatal(err)
	}

	loggers := logger.Loggers{}
	loggers.AjouterLogger(&leLogger)
	loggers.AjouterLogger(leLoggerF)

	loggers.Log("salut")

	// server := serveur.NewRestServerAgent("localhost:8080")
	// server.Start()
	// fmt.Scanln()
}
