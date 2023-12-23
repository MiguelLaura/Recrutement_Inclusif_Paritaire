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
	simu.etatInit = EtatSimulation{nbEmployes, pariteInit}
	simu.maxStep = maxStep

	simu.logger.AjouterLogger(logger.NewConsoleLogger())

	simu.ent = *NewEntreprise(nbEmployes, pariteInit, &simu.logger)
	recrut := NewRecrutement(&simu.ent, obj, sav, sap, trav, trap, ppav, ppap, &simu.logger)
	simu.ent.AjouterRecrutement(*recrut)

	simu.mettreAJourStatus(CREATED)

	return simu
}

// ---------------------
//        Getters
// ---------------------

// Pas de getters pour ent, logger et locker car les copies ne sont pas souhaitées (ils sont uniques)

func (simu *Simulation) PariteInit() float64 {
	return simu.pariteInit
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

func (simu *Simulation) EtatInit() EtatSimulation {
	return simu.etatInit
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

	simu.step = 0
	simu.mettreAJourStatus(STARTED)
	simu.start = time.Now()

	// Démarrage de l'entreprise
	go simu.ent.Start()

	simu.pariteInit = simu.ent.PourcentageFemmes()

	simu.locker.Add(1)
	go func() {
		for simu.step < simu.maxStep {
			if simu.status == STARTED {
				EnvoyerMessageEntreprise(&simu.ent, LIBRE, nil)
				simu.step++
				time.Sleep(2 * time.Second)
			} else if simu.status == PAUSED {
				time.Sleep(100 * time.Millisecond)
			} else if simu.status == ENDED {
				break
			}
		}

		// On s'assure que le statut de la simulation est bien à jour
		simu.mettreAJourStatus(ENDED)

		simu.terminerSimulation()
		simu.logger.Logf("La simulation est terminée.\nElle a duré : %v", time.Since(simu.start))
		simu.logger.LogType(LOG_REPONSE, ReponseAuClient{"stop", true})
		simu.locker.Done()
	}()

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

	simulationPrec := simu.ent.Recrutement()

	simu.ent = *NewEntreprise(
		simu.etatInit.nbEmp,
		simu.etatInit.parite,
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

	simu.logger.Log("La simulation a été récréé.")
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

	//création de la chaine de caractère pour le recrutement Avant
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
	if recrut.TypeRecrutementAvant() == PlacesReservees {
		infoRecrutementAvant = fmt.Sprintf("Places réservées : %d%%", int(recrut.PourcentagePlacesAvant()*100))
	}

	//création de la chaine de caractère pour le recrutement Après (si objectif)
	if recrut.Objectif() != -1 { // avec objectif
		infoRecrutementAvant = "(Avant) " + infoRecrutementAvant //ajout du qualificatif "avant" pour autre texte
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
			infoRecrutementApres = fmt.Sprintf("(Après) Compétences égales : %s", stratApres)
		}
		if recrut.TypeRecrutementApres() == PlacesReservees {

			infoRecrutementApres = fmt.Sprintf("(Après) Places réservées : %d%%", int(recrut.PourcentagePlacesApres()*100))
		}
	}

	simu.logger.LogType(LOG_INITIAL, InformationsInitiales{
		simu.PariteInit(), status, recrut.Objectif(), infoRecrutementAvant, infoRecrutementApres,
	})
}
