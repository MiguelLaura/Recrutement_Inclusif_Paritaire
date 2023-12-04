package agt

import (
	"fmt"
	"log"
	"time"
)

type Simulation struct {
	ent         Entreprise
	pariteInit  float32
	maxStep     int
	maxDuration time.Duration
	step        int // Stats
	start       time.Time
}

// retourne un pointeur sur une nouveau simulation
func NewSimulation(NbEmployes int, pariteInit float32, maxStep int, maxDuration time.Duration) (simu *Simulation) {
	simu = &Simulation{}
	simu.maxStep = maxStep
	simu.maxDuration = maxDuration

	simu.ent = *NewEntreprise(NbEmployes, pariteInit)

	return simu
}

// lance la simulation
func (simu *Simulation) Run() {
	log.Printf("Démarrage de la simulation [step: %d]", simu.step)

	// Démarrage du micro-service d'affichage
	go simu.Print()

	// On sauvegarde la date du début de la simulation
	simu.start = time.Now()
	pariteInit := simu.ent.PourcentageFemmes()

	// Démarrage de l'entreprise
	//simu.ent.Start() TO DO

	/*
		for {
			simu.step += 1
			simu.ent.Boucle ?
		}
	*/

	time.Sleep(simu.maxDuration)

	log.Printf("Fin de la simulation [step: %d, début parité : %f, fin parité : %f", simu.step, pariteInit, simu.ent.PourcentageFemmes())
}

// affiche les informations sur la parité au cours du temps
func (simu *Simulation) Print() {
	for {
		fmt.Printf("\rPourcentage de femmes = %f", simu.ent.PourcentageFemmes())
		time.Sleep(time.Second / 60) // 60 fps !
	}
}
