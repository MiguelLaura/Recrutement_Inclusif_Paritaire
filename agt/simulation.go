package agt

import (
	"log"
	"time"
)

// retourne un pointeur sur une nouvelle simulation
func NewSimulation(nbEmployes int, pariteInit float64, obj float64, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float64, ppap float64, maxStep int, maxDuration time.Duration) (simu *Simulation) {
	simu = &Simulation{}
	simu.maxStep = maxStep
	simu.maxDuration = maxDuration

	simu.ent = *NewEntreprise(nbEmployes, pariteInit)
	recrut := NewRecrutement(&simu.ent, obj, sav, sap, trav, trap, ppav, ppap)
	simu.ent.AjouterRecrutement(*recrut)

	simu.status = "created"

	return simu
}

func (simu *Simulation) Start() string {
	if simu.status != "created" {
		log.Println("La simulation ne peut pas être démarrée depuis cet état.")
		return "La simulation ne peut pas être démarrée depuis cet état."
	}

	simu.status = "started"
	simu.start = time.Now()

	// Démarrage de l'entreprise
	go simu.ent.Start()

	simu.pariteInit = simu.ent.PourcentageFemmes()

	msg_end := "La simulation démarre."

	go func() {
		for simu.step < simu.maxStep {
			if simu.status == "started" {
				EnvoyerMessageEntreprise(&simu.ent, LIBRE, nil)
				simu.step++
				time.Sleep(1 * time.Second)
			} else if simu.status == "pause" {
				time.Sleep(100 * time.Millisecond)
			} else if simu.status == "ended" {
				break
			}
		}

		log.Println("La simulation est terminée.")
		// Si on le récupère pas maintenant, les employés vont se terminer
		pariteFin := simu.ent.PourcentageFemmes()
		EnvoyerMessageEntreprise(&simu.ent, FIN, nil)
		// Permet d'attendre la fin effective de l'entreprise
		EnvoyerMessageEntreprise(&simu.ent, FIN, nil)

		log.Printf("Fin de la simulation [step: %d, nb employé fin : %d, début parité : %f, fin parité : %f]", simu.step, len(simu.ent.Employes()), simu.pariteInit, pariteFin)
	}()
	return msg_end

}

func (simu *Simulation) Pause() string {
	if simu.status != "started" {
		msg := "La simulation ne peut pas être mise en pause depuis cet état."
		log.Println(msg)
		return msg
	}
	simu.status = "pause"
	return "La simulation est en pause."
}

func (simu *Simulation) Continue() string {
	if simu.status != "pause" {
		msg := "La simulation ne peut pas être reprise depuis cet état."
		log.Println(msg)
		return msg
	}
	simu.status = "started"
	return "La simulation est relancée."
}

func (simu *Simulation) End() string {
	if simu.status == "ended" {
		msg := "La simulation est déjà terminée."
		log.Println(msg)
		return msg
	}
	simu.status = "ended"
	return "La simulation est terminée."
}
