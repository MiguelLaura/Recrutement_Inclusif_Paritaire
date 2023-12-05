package agt

import (
	"fmt"
	"log"
	"time"
)

// retourne un pointeur sur une nouvelle simulation
func NewSimulation(nbEmployes int, pariteInit float64, obj float64, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float64, ppap float64, maxStep int, maxDuration time.Duration) (simu *Simulation) {
	simu = &Simulation{}
	simu.maxStep = maxStep
	simu.maxDuration = maxDuration

	simu.ent = *NewEntreprise(nbEmployes, pariteInit)
	simu.pariteInit = simu.ent.PourcentageFemmes()
	recrut := NewRecrutement(&simu.ent, obj, sav, sap, trav, trap, ppav, ppap)
	simu.ent.AjouterRecrutement(*recrut)

	return simu
}

// lance la simulation
func (simu *Simulation) Run() {
	log.Printf("Démarrage de la simulation [step: %d]", simu.step)
	log.Printf("Nombre de steps à faire : %d", simu.maxStep)

	// Démarrage du micro-service d'affichage
	// go simu.Print()

	// On sauvegarde la date du début de la simulation
	simu.start = time.Now()

	// Démarrage de l'entreprise
	go simu.ent.Start()

	for simu.step < simu.maxStep {
		EnvoyerMessageEntreprise(&simu.ent, LIBRE, nil)
		simu.step += 1
	}
	EnvoyerMessageEntreprise(&simu.ent, FIN, nil)

	// time.Sleep(simu.maxDuration)

	log.Printf("Fin de la simulation [step: %d, début parité : %f, fin parité : %f", simu.step, simu.pariteInit, simu.ent.PourcentageFemmes())
}

// affiche les informations sur la parité au cours du temps
func (simu *Simulation) Print() {
	for {
		fmt.Printf("\rPourcentage de femmes = %f", simu.ent.PourcentageFemmes())
		time.Sleep(time.Second / 60) // 60 fps !
	}
}
