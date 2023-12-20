package agt

import (
	"log"
	"time"

	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

// retourne un pointeur sur une nouvelle simulation
func NewSimulation(nbEmployes int, pariteInit float64, obj float64, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float64, ppap float64, maxStep int) (simu *Simulation) {
	simu = &Simulation{}
	simu.maxStep = maxStep

	simu.logger.AjouterLogger(logger.NewConsoleLogger())

	simu.ent = *NewEntreprise(nbEmployes, pariteInit, &simu.logger)
	recrut := NewRecrutement(&simu.ent, obj, sav, sap, trav, trap, ppav, ppap, &simu.logger)
	simu.ent.AjouterRecrutement(*recrut)

	simu.status = CREATED

	return simu
}

func (simu *Simulation) Start() string {
	if simu.status != CREATED {
		log.Println("La simulation ne peut pas être démarrée depuis cet état.")
		return "La simulation ne peut pas être démarrée depuis cet état."
	}

	simu.status = STARTED
	simu.start = time.Now()

	// Démarrage de l'entreprise
	go simu.ent.Start()

	simu.pariteInit = simu.ent.PourcentageFemmes()

	msg_end := "La simulation démarre."

	go func() {
		for simu.step < simu.maxStep {
			if simu.status == STARTED {
				EnvoyerMessageEntreprise(&simu.ent, LIBRE, nil)
				simu.step++
				time.Sleep(1 * time.Second)
			} else if simu.status == PAUSED {
				time.Sleep(100 * time.Millisecond)
			} else if simu.status == ENDED {
				break
			}
		}

		log.Println("La simulation est terminée.")
		// Si on le récupère pas maintenant, les employés vont se terminer
		pariteFin := simu.ent.PourcentageFemmes()
		EnvoyerMessageEntreprise(&simu.ent, FIN, nil)
		// Permet d'attendre la fin effective de l'entreprise
		EnvoyerMessageEntreprise(&simu.ent, FIN, nil)

		log.Printf("Fin de la simulation [step: %d, nb employé fin : %d, début parité : %.2f, fin parité : %.2f]", simu.step, len(simu.ent.Employes()), simu.pariteInit, pariteFin)
	}()

	simu.logger.Log(msg_end)

	return msg_end

}

func (simu *Simulation) Pause() string {
	if simu.status != STARTED {
		msg := "La simulation ne peut pas être mise en pause depuis cet état."
		log.Println(msg)
		return msg
	}
	simu.status = PAUSED

	simu.logger.Log("La simulation est en pause.")

	return "La simulation est en pause."
}

func (simu *Simulation) Continue() string {
	if simu.status != PAUSED {
		msg := "La simulation ne peut pas être reprise depuis cet état."
		log.Println(msg)
		return msg
	}
	simu.status = STARTED

	simu.logger.Log("La simulation est relancée.")

	return "La simulation est relancée."
}

func (simu *Simulation) End() string {
	if simu.status == ENDED {
		msg := "La simulation est déjà terminée."
		log.Println(msg)
		return msg
	}
	simu.status = ENDED

	simu.logger.Logf("La simulation est terminée.\nElle a duré : %v", time.Since(simu.start))

	return "La simulation est terminée."
}

func (simu *Simulation) AjouteWebSockerLogger(wsLogger *logger.SocketLogger) {
	simu.logger.AjouterLogger(wsLogger)
}
