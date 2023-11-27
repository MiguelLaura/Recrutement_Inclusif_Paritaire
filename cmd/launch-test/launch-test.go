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

	loggers.Log("cc ", "comment ", "ça ", "va ", "?")
}
