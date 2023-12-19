package main

import (
	"fmt"

	"gitlab.utc.fr/mennynat/ia04-project/serveur"
)

func main() {

	// leLogger := logger.ConsoleLogger{}
	// leLoggerF, err := logger.NewFileLogger("log/myLog/logFile/")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// loggers := logger.Loggers{}
	// loggers.AjouterLogger(&leLogger)
	// loggers.AjouterLogger(leLoggerF)

	// loggers.Log("salut", "comment", "ça", "va")
	// loggers.LogType("MonType de log", "salut", "au revoir")

	server := serveur.NewRestServerAgent("localhost:8080")
	server.Start()
	fmt.Scanln()
}
