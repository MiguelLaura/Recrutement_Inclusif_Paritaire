package main

import (
	"gitlab.utc.fr/mennynat/ia04-project/agt"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

func main() {

	leLogger := logger.ConsoleLogger{}

	loggers := logger.Loggers{}
	loggers.AjouterLogger(&leLogger)

	//loggers.Log("cc", "comment", "ça", "va", "?")

	//leLogger.Log("cc", "comment", "ça", "va", "?")

	var e agt.Entreprise
	leLogger.Log(agt.GenererCandidats(10, e))

}
