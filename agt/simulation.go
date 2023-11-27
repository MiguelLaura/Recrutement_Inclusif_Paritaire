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

func NewSimulation(NbEmployes int, pariteInit float32, maxStep int, maxDuration time.Duration) (simu *Simulation) {
	simu = &Simulation{}
	simu.maxStep = maxStep
	simu.maxDuration = maxDuration

	//simu.ent = *NewEntreprise(NbEmployes, pariteInit)

	return simu
}

func (simu *Simulation) Run() {
	log.Printf("Démarrage de la simulation [step: %d]", simu.step)

	// Démarrage du micro-service de Log
	//go simu.Log()
	// Démarrage du micro-service d'affichage
	//go simu.Print()

	// On sauvegarde la date du début de la simulation
	simu.start = time.Now()

	// Démarrage de l'entreprise
	//simu.ent.Start() TO DO

	time.Sleep(simu.maxDuration)

	//log.Printf("Fin de la simulation [step: %d, in: %d, out: %d, π: %f]", simu.step, simu.env.in, simu.env.out, simu.env.PI())
}

func (simu *Simulation) Print() {
	for {
		fmt.Printf("\rparite = %.30f" /*simu.ent.NBFemmes*/) //A VOIR
		time.Sleep(time.Second / 60)                         // 60 fps !
	}
}

func (simu *Simulation) Log() {
	// Not implemented
}
