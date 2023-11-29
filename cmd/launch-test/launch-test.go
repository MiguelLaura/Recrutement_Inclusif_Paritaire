package main

import (
	"gitlab.utc.fr/mennynat/ia04-project/agt"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

func main() {

	leLogger := logger.ConsoleLogger{}

	loggers := logger.Loggers{}
	loggers.AjouterLogger(&leLogger)

	var recrutnil agt.Recrutement
	e := agt.NewEntreprise(5, 0.5, recrutnil)
	r := agt.NewRecrutement(e, -1, agt.PrioFemme, agt.StratVide, agt.Competences, agt.Vide, -1, -1)
	embauches, err := r.GenererCandidats(10)
	leLogger.Log("candidats: ", embauches)
	leLogger.Log(err)

}
