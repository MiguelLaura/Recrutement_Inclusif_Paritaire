package main

import (
	"fmt"

	"gitlab.utc.fr/mennynat/ia04-project/serveur"
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

	server := serveur.NewRestServerAgent("localhost:8080")
	server.Start()
	fmt.Scanln()
}
