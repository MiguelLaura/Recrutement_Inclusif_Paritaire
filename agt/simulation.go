package agt

import (
	"fmt"
	"log"
	"time"

	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

// ---------------------
//     Constructeur
// ---------------------

// retourne un pointeur sur une nouvelle simulation
func NewSimulation(nbEmployes int, pariteInit float64, obj float64, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float64, ppap float64, maxStep int) (simu *Simulation) {
	simu = &Simulation{}
	simu.maxStep = maxStep
	simu.agentsLances = false

	simu.logger.AjouterLogger(logger.NewConsoleLogger())

	simu.ent = *NewEntreprise(nbEmployes, pariteInit, &simu.logger)
	recrut := NewRecrutement(&simu.ent, obj, sav, sap, trav, trap, ppav, ppap, &simu.logger)
	simu.ent.AjouterRecrutement(*recrut)

	simu.mettreAJourStatus(CREATED)

	simu.pariteInit = pariteInit
	simu.nbEmployesInit = nbEmployes

	return simu
}

// ---------------------
//        Getters
// ---------------------

// Pas de getters pour ent, logger et locker car les copies ne sont pas souhaitées (ils sont uniques)

func (simu *Simulation) PariteInit() float64 {
	return simu.pariteInit
}

func (simu *Simulation) NbEmployeInit() int {
	return simu.nbEmployesInit
}

func (simu *Simulation) MaxStep() int {
	return simu.maxStep
}

func (simu *Simulation) Step() int {
	return simu.step
}

func (simu *Simulation) StartTime() time.Time {
	return simu.start
}

func (simu *Simulation) Status() Status {
	return simu.status
}

func (simu *Simulation) AgentsLances() bool {
	return simu.agentsLances
}

// ---------------------
//        Setters
// ---------------------

func (simu *Simulation) SetPariteInit(pariteInit float64) {
	simu.pariteInit = pariteInit
}

func (simu *Simulation) SetnbEmployesInit(nbEmployesInit int) {
	simu.nbEmployesInit = nbEmployesInit
}

func (simu *Simulation) SetMaxStep(maxStep int) {
	simu.maxStep = maxStep
}

func (simu *Simulation) SetStep(step int) {
	simu.step = step
}

func (simu *Simulation) SetStartTime(startTime time.Time) {
	simu.start = startTime
}

func (simu *Simulation) SetStatus(status Status) {
	simu.status = status
}

func (simu *Simulation) SetAgentsLances(lances bool) {
	simu.agentsLances = lances
}

// ---------------------
//  Logique de simulation
// ---------------------

func (simu *Simulation) Start() {
	if simu.status != CREATED {
		simu.logger.Err("La simulation ne peut pas être démarrée depuis cet état.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"start", false})
		return
	}

	simu.startAgents()
	simu.mettreAJourStatus(STARTED)

	simu.logger.Log("La simulation démarre.")
	simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"start", true})
}

func (simu *Simulation) Pause() {
	if simu.status != STARTED {
		simu.logger.Err("La simulation ne peut pas être mise en pause depuis cet état.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"pause", false})
		return
	}

	simu.mettreAJourStatus(PAUSED)

	simu.logger.Log("La simulation est en pause.")
	simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"pause", true})
}

func (simu *Simulation) Continue() {
	if simu.status != PAUSED {
		simu.logger.Err("La simulation ne peut pas être reprise depuis cet état.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"continue", false})
		return
	}

	simu.mettreAJourStatus(STARTED)

	simu.logger.Log("La simulation est relancée.")
	simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"continue", true})
}

func (simu *Simulation) NextStep() {
	if simu.status == STARTED || simu.status == ENDED {
		simu.logger.Err("La simulation ne peut pas continuer ainsi.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"step", false})
		return
	}

	if simu.status == CREATED {
		simu.startAgents()
	}

	simu.mettreAJourStatus(STEP)

	simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"step", true})
}

func (simu *Simulation) End() {
	if simu.status == ENDED {
		simu.logger.Err("La simulation est déjà terminée.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"stop", false})
		return
	}

	if simu.status == CREATED {
		simu.logger.Err("La simulation n'a pas encore commencé.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"stop", false})
		return
	}

	simu.mettreAJourStatus(ENDED)

	simu.locker.Wait()

	//informations sur la fin (log & durée dans la boucle de simulation)

}

func (simu *Simulation) Relancer() {
	if simu.status == CREATED {
		simu.logger.Err("La simulation n'a pas encore commencé.")
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"relancer", false})
		return
	}

	if simu.status != ENDED {
		log.Printf("Termine la simulation en cours")
		simu.mettreAJourStatus(ENDED)
		simu.locker.Wait()
	}

	simu.agentsLances = false
	simulationPrec := simu.ent.Recrutement()

	simu.ent = *NewEntreprise(
		simu.nbEmployesInit,
		simu.pariteInit,
		&simu.logger,
	)

	recrut := NewRecrutement(
		&simu.ent,
		simulationPrec.Objectif(),
		simulationPrec.StratAvant(),
		simulationPrec.StratApres(),
		simulationPrec.TypeRecrutementAvant(),
		simulationPrec.TypeRecrutementApres(),
		simulationPrec.PourcentagePlacesAvant(),
		simulationPrec.PourcentagePlacesApres(),
		&simu.logger,
	)

	simu.ent.AjouterRecrutement(*recrut)
	simu.mettreAJourStatus(CREATED)

	simu.logger.Log("La simulation a été recréée.")
	simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"relancer", true})
}

func (simu *Simulation) AjouteWebSockerLogger(wsLogger *logger.SocketLogger) {
	simu.logger.AjouterLogger(wsLogger)
}

func (simu *Simulation) terminerSimulation() {
	if simu.status != ENDED {
		return
	}

	// Si on le récupère pas maintenant, les employés vont se terminer
	pariteFin := simu.ent.PourcentageFemmes()
	EnvoyerMessageEntreprise(&simu.ent, FIN, nil)
	// Permet d'attendre la fin effective de l'entreprise
	EnvoyerMessageEntreprise(&simu.ent, FIN, nil)

	log.Printf("Fin de la simulation [step: %d, nb employé fin : %d, début parité : %.2f, fin parité : %.2f]", simu.step, len(simu.ent.Employes()), simu.pariteInit, pariteFin)
}

func (simu *Simulation) mettreAJourStatus(nouveauStatus Status) {
	simu.locker.Lock()
	defer simu.locker.Unlock()

	simu.status = nouveauStatus
}

// -------------------------------------
//  Fonctions pour envoyer informations
// -------------------------------------

func (simu *Simulation) EnvoyerInfosInitiales() {
	status := ""

	if simu.status == CREATED {
		status = "not_started"
	}
	if simu.status == STARTED {
		status = "start"
	}
	if simu.status == PAUSED {
		status = "pause"
	}
	if simu.status == ENDED {
		status = "stop"
	}

	recrut := simu.ent.Recrutement()
	infoRecrutementAvant := ""
	infoRecrutementApres := ""

	//création de la chaine de caractère pour le recrutement en dessous du seuil
	if recrut.TypeRecrutementAvant() == Competences {
		stratAvant := ""
		if recrut.StratAvant() == PrioFemme {
			stratAvant = "Femmes"
		}
		if recrut.StratAvant() == PrioHomme {
			stratAvant = "Hommes"
		}
		if recrut.StratAvant() == Hasard {
			stratAvant = "Hasard"
		}
		infoRecrutementAvant = fmt.Sprintf("Compétences égales : %s", stratAvant)
	}
	if recrut.TypeRecrutementAvant() == PlacesReserveesFemme {
		infoRecrutementAvant = fmt.Sprintf("Places réservées aux femmes : %d%%", int(recrut.PourcentagePlacesAvant()*100))
	}
	if recrut.TypeRecrutementAvant() == PlacesReserveesHomme {
		infoRecrutementAvant = fmt.Sprintf("Places réservées aux hommes : %d%%", int(recrut.PourcentagePlacesAvant()*100))
	}

	//création de la chaine de caractère pour le recrutement au dessus du seuil (si objectif)
	if recrut.Objectif() != -1 { // avec objectif
		infoRecrutementAvant = "(En dessous du seuil) " + infoRecrutementAvant //ajout du qualificatif "en dessous du seuil" pour autre texte
		if recrut.TypeRecrutementApres() == Competences {
			stratApres := ""
			if recrut.StratApres() == PrioFemme {
				stratApres = "Femmes"
			}
			if recrut.StratApres() == PrioHomme {
				stratApres = "Hommes"
			}
			if recrut.StratApres() == Hasard {
				stratApres = "Hasard"
			}
			infoRecrutementApres = fmt.Sprintf("(Au dessus du seuil) Compétences égales : %s", stratApres)
		}
		if recrut.TypeRecrutementApres() == PlacesReserveesFemme {
			infoRecrutementApres = fmt.Sprintf("(Au dessus du seuil) Places réservées aux femmes : %d%%", int(recrut.PourcentagePlacesApres()*100))
		}
		if recrut.TypeRecrutementApres() == PlacesReserveesHomme {
			infoRecrutementApres = fmt.Sprintf("(Au dessus du seuil) Places réservées aux hommes : %d%%", int(recrut.PourcentagePlacesApres()*100))
		}
	}

	simu.logger.LogType(LOG_INITIAL, InformationsInitiales{
		simu.maxStep, simu.PariteInit(), simu.NbEmployeInit(), status, recrut.Objectif(), infoRecrutementAvant, infoRecrutementApres,
	})
}

func (simu *Simulation) obtenirSituationActuelle() SituationActuelle {

	nbemp := simu.ent.NbEmployes()
	parite := simu.ent.PourcentageFemmes()
	benef := simu.ent.CalculerBenefice()
	competence := simu.ent.MoyenneCompetences()
	santeMentale := simu.ent.MoyenneSanteMentale()
	situ := NewSituationActuelle(simu.step, nbemp, parite, benef, competence, santeMentale)
	return *situ

}

func (simu *Simulation) startAgents() {
	simu.locker.Lock()
	defer simu.locker.Unlock()

	if simu.agentsLances {
		return
	}

	simu.step = 0
	simu.start = time.Now()

	go simu.ent.Start()

	simu.locker.Add(1)

	go func() {

	BOUCLE_SIMULATION:
		for simu.step < simu.maxStep {
			switch simu.status {
			case STARTED, STEP:
				EnvoyerMessageEntreprise(&simu.ent, LIBRE, nil)
				simu.step++

				if simu.status == STEP {
					simu.mettreAJourStatus(PAUSED)
				}

				simu.ent.logger.LogType(LOG_GLOBAL, simu.obtenirSituationActuelle())

				time.Sleep(2 * time.Second)
			case PAUSED:
				time.Sleep(100 * time.Millisecond)
			case ENDED:
				break BOUCLE_SIMULATION
			}
		}

		// On s'assure que le statut de la simulation est bien à jour
		simu.mettreAJourStatus(ENDED)

		simu.terminerSimulation()
		simu.logger.Logf("La simulation est terminée.\nElle a duré : %v", time.Since(simu.start).Round(time.Second))
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"stop", true})
		simu.locker.Done()
	}()

	simu.agentsLances = true
}
