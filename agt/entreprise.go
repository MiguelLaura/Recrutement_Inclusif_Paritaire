package agt

import (
	"log"
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

// ---------------------
//        Général
// ---------------------

func EnvoyerMessageEntreprise(dest *Entreprise, act Action, payload any) {
	dest.chnl <- Communicateur{act, payload}
}

func EnvoyerNotifActions(dest *Entreprise, act Action, payload any) {
	dest.chnlNotifAction <- Communicateur{act, payload}
}

func EnvoyerMessageRecrutement(dest *Recrutement, act ActionRecrutement, payload any) {
	dest.Chnl() <- CommunicateurRecrutement{act, payload}
}

// ---------------------
//     Constructeur
// ---------------------

func NewCompteur() *Compteur {
	return &Compteur{
		nbEmbauches:           0,
		nbEmbauchesFemme:      0,
		nbAgressions:          0,
		nbPlaintes:            0,
		nbDemissions:          0,
		nbDemissionsMaternite: 0,
		nbRetraites:           0,
		nbLicenciements:       0,
		nbDepressions:         0,
		nbEnfants:             0,
		nbCongesPaternite:     0,
		nbCongesMaternite:     0,
	}
}

// La fonction NewEntreprise doit créer l'entreprise et générer les employés de façon à respecter le quota de parité initial
func NewEntreprise(nbEmployesInit int, pariteInit float64, logger *logger.Loggers) *Entreprise {
	ent := new(Entreprise)

	var nbFemmes int = int(math.Round(float64(nbEmployesInit) * float64(pariteInit)))
	var nbHommes int = nbEmployesInit - nbFemmes
	var employesInit []*Employe

	for i := 0; i < nbFemmes; i++ {
		emp := GenererEmployeInit(&ent, Femme, logger)
		employesInit = append(employesInit, emp)
		if emp.Agresseur() {
			ent.nbAgresseurs += 1
		}
	}
	for i := 0; i < nbHommes; i++ {
		emp := GenererEmployeInit(&ent, Homme, logger)
		employesInit = append(employesInit, emp)
		if emp.Agresseur() {
			ent.nbAgresseurs += 1
		}
	}
	ent.employes = employesInit
	departs := make([]*Employe, 0)
	ent.departs = departs
	formation := make([]*Employe, 0)
	ent.formation = formation
	congeParental := make([]*Employe, 0)
	ent.congeParental = congeParental
	plaintes := make([][]*Employe, 0)
	ent.plaintes = plaintes
	ent.cmpt = NewCompteur()
	ent.nbActions = 0
	ent.fin = false
	ent.chnl = make(chan Communicateur)
	ent.chnlRecrutement = make(chan CommunicateurRecrutement)
	ent.chnlNotifAction = make(chan Communicateur)
	ent.logger = logger
	return ent
}

// ---------------------
//        Getters
// ---------------------

func (ent *Entreprise) Employes() []*Employe {
	return ent.employes
}

func (ent *Entreprise) Departs() []*Employe {
	return ent.departs
}

func (ent *Entreprise) Formation() []*Employe {
	return ent.formation
}

func (ent *Entreprise) CongeParental() []*Employe {
	return ent.congeParental
}

func (ent *Entreprise) Plaintes() [][]*Employe {
	return ent.plaintes
}

func (ent *Entreprise) Cmpt() *Compteur {
	return ent.cmpt
}

func (ent *Entreprise) NbEmbauches() int {
	return ent.cmpt.nbEmbauches
}

func (ent *Entreprise) NbEmbauchesFemme() int {
	return ent.cmpt.nbEmbauchesFemme
}

func (ent *Entreprise) NbAgressions() int {
	return ent.cmpt.nbAgressions
}

func (ent *Entreprise) NbPlaintes() int {
	return ent.cmpt.nbPlaintes
}

func (ent *Entreprise) NbDemissions() int {
	return ent.cmpt.nbDemissions
}

func (ent *Entreprise) NbDemissionsMaternite() int {
	return ent.cmpt.nbDemissionsMaternite
}

func (ent *Entreprise) NbRetraites() int {
	return ent.cmpt.nbRetraites
}

func (ent *Entreprise) NbLicenciements() int {
	return ent.cmpt.nbLicenciements
}

func (ent *Entreprise) NbDepressions() int {
	return ent.cmpt.nbDepressions
}

func (ent *Entreprise) NbDeparts() int {
	return ent.NbLicenciements() + ent.NbDepressions() + ent.NbDemissions() + ent.NbDemissionsMaternite() + ent.NbRetraites()
}

func (ent *Entreprise) NbEnfants() int {
	ent.RLock()
	defer ent.RUnlock()
	return ent.cmpt.nbEnfants
}

func (ent *Entreprise) NbCongesMaternite() int {
	return ent.cmpt.nbCongesMaternite
}

func (ent *Entreprise) NbCongesPaternite() int {
	return ent.cmpt.nbCongesPaternite
}

func (ent *Entreprise) Recrutement() Recrutement {
	return ent.recrutement
}

func (ent *Entreprise) NbActions() int {
	return ent.nbActions
}

func (ent *Entreprise) NbAgresseurs() int {
	return ent.nbAgresseurs
}

func (ent *Entreprise) Fin() bool {
	return ent.fin
}

func (ent *Entreprise) Chnl() chan Communicateur {
	return ent.chnl
}

func (ent *Entreprise) ChnlRecrutement() chan CommunicateurRecrutement {
	return ent.chnlRecrutement
}

func (ent *Entreprise) ChnlNotifAction() chan Communicateur {
	return ent.chnlNotifAction
}

func (ent *Entreprise) Logger() *logger.Loggers {
	return ent.logger
}

// ---------------------
//        Setters
// ---------------------

func (ent *Entreprise) SetEmployes(employes []*Employe) {
	ent.employes = employes
}

func (ent *Entreprise) SetDeparts(departs []*Employe) {
	ent.departs = departs
}

func (ent *Entreprise) SetFormation(formation []*Employe) {
	ent.formation = formation
}

func (ent *Entreprise) SetCongeParental(congeParental []*Employe) {
	ent.congeParental = congeParental
}

func (ent *Entreprise) SetPlaintes(plaintes [][]*Employe) {
	ent.plaintes = plaintes
}

func (ent *Entreprise) SetCompteur(cmpt *Compteur) {
	ent.cmpt = cmpt
}

func (ent *Entreprise) SetNbEmbauches(nbEmbauches int) {
	ent.cmpt.nbEmbauches = nbEmbauches
}

func (ent *Entreprise) SetNbEmbauchesFemme(nbEmbauchesFemme int) {
	ent.cmpt.nbEmbauchesFemme = nbEmbauchesFemme
}

func (ent *Entreprise) SetNbAgressions(nbAgressions int) {
	ent.cmpt.nbAgressions = nbAgressions
}

func (ent *Entreprise) SetNbPlaintes(nbPlaintes int) {
	ent.cmpt.nbPlaintes = nbPlaintes
}

func (ent *Entreprise) SetNbDemissions(nbDemissions int) {
	ent.cmpt.nbDemissions = nbDemissions
}

func (ent *Entreprise) SetNbDemissionsMaternite(nbDemissionsMaternite int) {
	ent.cmpt.nbDemissionsMaternite = nbDemissionsMaternite
}

func (ent *Entreprise) SetNbRetraites(nbRetraites int) {
	ent.cmpt.nbRetraites = nbRetraites
}

func (ent *Entreprise) SetNbLicenciements(nbLicenciements int) {
	ent.cmpt.nbLicenciements = nbLicenciements
}

func (ent *Entreprise) SetNbDepressions(nbDepressions int) {
	ent.cmpt.nbDepressions = nbDepressions
}

func (ent *Entreprise) SetNbEnfants(nbEnfants int) {
	ent.Lock()
	defer ent.Unlock()
	ent.cmpt.nbEnfants = nbEnfants
}

func (ent *Entreprise) SetNbCongesMaternite(nbCongesMaternite int) {
	ent.cmpt.nbCongesMaternite = nbCongesMaternite
}

func (ent *Entreprise) SetNbCongesPaternite(nbCongesPaternite int) {
	ent.cmpt.nbCongesPaternite = nbCongesPaternite
}

func (ent *Entreprise) SetRecrutement(recrut Recrutement) {
	ent.recrutement = recrut
}

func (ent *Entreprise) SetNbActions(nbActions int) {
	ent.nbActions = nbActions
}

func (ent *Entreprise) SetNbAgresseurs(nbAgresseurs int) {
	ent.nbAgresseurs = nbAgresseurs
}

func (ent *Entreprise) SetFin(fin bool) {
	ent.fin = fin
}

func (ent *Entreprise) SetChnl(chnl chan Communicateur) {
	ent.chnl = chnl
}

func (ent *Entreprise) SetChnlRecrutement(chnlRecrutement chan CommunicateurRecrutement) {
	ent.chnlRecrutement = chnlRecrutement
}

func (ent *Entreprise) SetChnlNotifAction(chnlNotifAction chan Communicateur) {
	ent.chnlNotifAction = chnlNotifAction
}

func (ent *Entreprise) SetLogger(logger *logger.Loggers) {
	ent.logger = logger
}

// ---------------------
//     En cours d'année (appelées par les employés)
// ---------------------

func (ent *Entreprise) RecevoirDemission(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()
	ent.SetNbDemissions(ent.NbDemissions() + 1)
	i, _ := TrouverEmploye(ent.departs, func(e *Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		ent.departs = append(ent.departs, emp)
		log.Printf("%s pose sa démission", emp.String())
	}
}

func (ent *Entreprise) RecevoirDemissionMaternite(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	i, _ := TrouverEmploye(ent.departs, func(e *Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		ent.SetNbDemissionsMaternite(ent.NbDemissionsMaternite() + 1)
		ent.departs = append(ent.departs, emp)
		log.Printf("%s pose sa démission après son congé maternité", emp.String())

	}
}

func (ent *Entreprise) RecevoirDepression(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	i, _ := TrouverEmploye(ent.departs, func(e *Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		ent.SetNbDepressions(ent.NbDepressions() + 1)
		ent.departs = append(ent.departs, emp)
		log.Printf("%s pose sa démission pour dépression", emp.String())

	}
}

func (ent *Entreprise) RecevoirRetraite(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	i, _ := TrouverEmploye(ent.departs, func(e *Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		ent.SetNbRetraites(ent.NbRetraites() + 1)
		ent.departs = append(ent.departs, emp)
		log.Printf("%s part à la retraite", emp.String())

	}
}

func (ent *Entreprise) RecevoirCongeParental(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	i, _ := TrouverEmploye(ent.congeParental, func(e *Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		if emp.Genre() == Femme {
			ent.SetNbCongesMaternite(ent.NbCongesMaternite() + 1)
		} else {
			ent.SetNbCongesPaternite(ent.NbCongesPaternite() + 1)
		}
		ent.congeParental = append(ent.congeParental, emp)
		log.Printf("%s part en congé parental", emp.String())
	}
}

func (ent *Entreprise) RecevoirPlainte(plaignant *Employe, accuse *Employe) {
	ent.Lock()
	defer ent.Unlock()
	ent.SetNbPlaintes(ent.NbPlaintes() + 1)
	ent.plaintes = append(ent.plaintes, []*Employe{plaignant, accuse})
	log.Printf("%s porte plainte contre %s ", plaignant.String(), accuse.String())
}

func (ent *Entreprise) RecevoirActions(nbActions int) {
	for {
		<-ent.chnlNotifAction

		ent.nbActions += 1

		if ent.nbActions == nbActions {
			ent.nbActions = 0
			return
		}
	}
}

// ---------------------
//     En cours d'année (appelées par l'entreprise)
// ---------------------

func (ent *Entreprise) teamBuilding() {
	ent.logger.LogType(LOG_EVENEMENT, "Un team building est organisé. Les employé·e·s sont content·e·s.")
	for _, e := range ent.employes {
		testPresence, _ := TrouverEmploye(ent.departs, func(emp *Employe) bool { return e.Id() == emp.Id() }, 0)
		// On vérifie que l'employé ne va pas partir
		if testPresence < 0 && (e.SanteMentale() < 100) {
			if (e.SanteMentale() + constantes.BOOST_TEAM_BUILDING) > 100 {
				e.SetSanteMentale(100)
			} else {
				e.SetSanteMentale(e.SanteMentale() + constantes.BOOST_TEAM_BUILDING)
			}
		}
	}
}

func (ent *Entreprise) organisationFormation() {

	ent.formation = make([]*Employe, 0)
	// Génération des employés participant à une formation cette année

	// 32% des français ont participé à une formation
	nb_employes_formes := math.Round(constantes.POURCENTAGE_FORMATION * float64(ent.NbEmployes()))
	if nb_employes_formes != 0 {
		ent.logger.LogfType(LOG_EVENEMENT, "%d employé·e(s) ont participé à une formation.", int(nb_employes_formes))
		// 50% des employés qui se forment sont des femmes
		nb_femmes_formes := math.Round(nb_employes_formes / 2)
		nb_hommes_formes := nb_femmes_formes
		femmes := FiltreFemme(ent.employes)
		hommes := FiltreHomme(ent.employes)
		for idx := 0; idx < int(nb_femmes_formes); idx++ {
			if len(femmes) == 0 {
				break
			}
			i := rand.Intn(len(femmes))
			ent.formation = append(ent.formation, femmes[i])
			// Pour ne pas avoir de doublons
			femmes = enleverEmploye(femmes, femmes[i])
		}
		for idx := 0; idx < int(nb_hommes_formes); idx++ {
			if len(hommes) == 0 {
				break
			}
			i := rand.Intn(len(hommes))
			ent.formation = append(ent.formation, hommes[i])
			// Pour ne pas avoir de doublons
			hommes = enleverEmploye(hommes, hommes[i])
		}
	}

}

// ---------------------
//     Fin d'année
// ---------------------

// // Renvoyer selon un certain pourcentage
func (ent *Entreprise) gestionPlaintes() {
	if len(ent.plaintes) <= 0 {
		return
	}
	for _, e := range ent.plaintes {
		if rand.Float64() <= constantes.PROBA_LICENCIEMENT {
			accuse := e[1]
			i, _ := TrouverEmploye(ent.departs, func(e *Employe) bool { return e.Id() == accuse.Id() }, 0)
			if i < 0 {
				ent.SetNbLicenciements(ent.NbLicenciements() + 1)
				ent.departs = append(ent.departs, accuse)
				log.Printf("%s est licencié·e pour faute grave", accuse.String())
			}
		}
	}
	ent.plaintes = make([][]*Employe, 0)
}

func (ent *Entreprise) gestionDeparts() {
	if len(ent.departs) <= 0 {
		return
	}
	for _, emp := range ent.departs {
		ent.employes = enleverEmploye(ent.employes, emp)
		go func(emp *Employe) {
			EnvoyerMessage(emp, FIN, nil)
		}(emp)
		if emp.Agresseur() {
			ent.nbAgresseurs -= 1
		}
	}
	ent.RecevoirActions(len(ent.departs))
	ent.departs = make([]*Employe, 0)
}

func (ent *Entreprise) gestionRecrutements() {
	msg := <-ent.chnlRecrutement
	if msg.Act == FIN_RECRUTEMENT {
		embauches := msg.Payload.([]*Employe)
		for _, emp := range embauches {
			ent.SetNbEmbauches(ent.NbEmbauches() + 1)
			if emp.Genre() == Femme {
				ent.SetNbEmbauchesFemme(ent.NbEmbauchesFemme() + 1)
			}
			ent.employes = append(ent.employes, emp)
			log.Printf("%s est embauché·e", emp.String())
			if emp.Agresseur() {
				ent.nbAgresseurs += 1
			}
		}
	}
}

func (ent *Entreprise) lancerRecrutements() {
	nbARecruter := float64(ent.NbEmployes())*constantes.POURCENTAGE_RECRUTEMENT + 1.0

	go EnvoyerMessageRecrutement(&ent.recrutement, RECRUTEMENT, int(nbARecruter))
}

func (ent *Entreprise) CalculerBenefice() int {
	var benef float64 = 0

	// Pour chaque employé, on calcule ce qu'il rapporte à l'entreprise en fonction de sa santé mentale et compétences
	// Compétences varient entre 0 et 10 par une loi normale d'espérance 5.
	// CA_PAR_EMPLOYE/5 * competences pour garder la valeur moyenne du CA_PAR_EMPLOYE mais prendre en compte les compétences
	// CA_PAR_EMPLOYE dépend de la taille de l'entreprise
	// benef plus faible si santé mentale plus basse
	if ent.NbEmployes() < 10 {
		for _, e := range ent.employes {
			benef += (constantes.CA_PAR_EMPLOYE_MIC/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE
		}
		// Impact conges parental : pas de salaire et pas de productivité pendant une certaine période
		for _, e := range ent.congeParental {
			if e.Genre() == Femme {
				benef -= constantes.PROPORTION_ARRET_F * ((constantes.CA_PAR_EMPLOYE_MIC/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE)
			} else {
				benef -= constantes.PROPORTION_ARRET_H * ((constantes.CA_PAR_EMPLOYE_MIC/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE)
			}
		}
	} else if ent.NbEmployes() >= 10 && ent.NbEmployes() < 250 {
		for _, e := range ent.employes {
			benef += (constantes.CA_PAR_EMPLOYE_PME/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE
		}

		// Impact conges parental : pas de salaire et pas de productivité pendant une certaine période
		for _, e := range ent.congeParental {
			if e.Genre() == Femme {
				benef -= constantes.PROPORTION_ARRET_F * ((constantes.CA_PAR_EMPLOYE_PME/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE)
			} else {
				benef -= constantes.PROPORTION_ARRET_H * ((constantes.CA_PAR_EMPLOYE_PME/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE)
			}
		}
	} else {
		for _, e := range ent.employes {
			benef += (constantes.CA_PAR_EMPLOYE_ETI/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE
		}
		// Impact conges parental : pas de salaire et pas de productivité pendant une certaine période
		for _, e := range ent.congeParental {
			if e.Genre() == Femme {
				benef -= constantes.PROPORTION_ARRET_F * ((constantes.CA_PAR_EMPLOYE_ETI/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE)
			} else {
				benef -= constantes.PROPORTION_ARRET_H * ((constantes.CA_PAR_EMPLOYE_ETI/5)*float64(e.Competence())*float64(e.SanteMentale())/100 - constantes.COUT_EMPLOYE)
			}
		}
	}

	// Bonus de productivité si %femmes supérieur à 35%
	if ent.PourcentageFemmes() > constantes.SEUIL_IMPACT_FEMME {
		benef = benef * (1.0 + constantes.BOOST_PRODUCTIVITE_FEMME)
		ent.logger.LogfType(LOG_ENTREPRISE, "Il y a plus de 35%% de femmes dans l'entreprise, ce qui a permis une ambiance productive : amélioration du bénéfice.")
	}

	// Coût du recrutement
	nbARecruter := float64(ent.NbEmployes())*constantes.POURCENTAGE_RECRUTEMENT + 1.0
	benef -= float64(nbARecruter * constantes.COUT_RECRUTEMENT)

	// Coût du teambuilding
	benef -= 2 * float64(ent.NbEmployes()) * constantes.COUT_TB_PAR_EMPLOYE

	// Coût des formations
	nbFormes := len(ent.formation)
	benef -= float64(constantes.PRIX_FORMATION * constantes.NB_JOURS_FORMATION * nbFormes)

	// Amende si non parité
	// Modèle le plus simple : si %Femmes ne respectent pas la loi (<40%), amende d'1% des bénéfices
	// Modèle 2 plus proche de la réalité : amende si non respect pendant 3 ans consécutifs
	// Modèle 3 le plus réaliste : amende à partir de 2029
	if ent.PourcentageFemmes() < constantes.SEUIL_AMENDE {
		amende := benef * constantes.POURCENTAGE_AMENDE
		ent.logger.LogfType(LOG_ENTREPRISE, "L'entreprise ne respecte pas la loi Rixain sur la parité (40%% de femmes minimum) et doit payer une amende de %d euros.", int(math.Round(amende)))
		benef -= amende
	}

	return int(math.Round(benef))
}

func (ent *Entreprise) bonneAnnee() {
	ent.resetCompteur()
	ent.SetNbAgressions(ent.NbAgresseurs())
	ent.congeParental = make([]*Employe, 0)

	for _, emp := range ent.employes {
		go func(emp *Employe) {
			EnvoyerMessage(emp, LIBRE, nil)
		}(emp)
	}

	ent.lancerRecrutements()
}

// ---------------------
//  Logique de simulation
// ---------------------

func (ent *Entreprise) Start() {
	go func() {
		for _, emp := range ent.employes {
			emp.Start()
		}

		ent.recrutement.Start()

		for {
			msg := <-ent.chnl
			if msg.Act == LIBRE && !ent.fin {
				ent.agir()
			} else if msg.Act == FIN && !ent.fin {
				ent.stop()
				break
			} else {
				msg = <-ent.chnl
				if msg.Act == FIN {
					break
				}
			}
		}
		log.Printf("Fin d'entreprise")
		<-ent.chnl
	}()
}

func (ent *Entreprise) agir() {
	if ent.NbEmployes() <= 0 {
		ent.fin = true
		return
	}
	ent.logger.LogType(LOG_ENTREPRISE, "Début d'année.")
	log.Printf("Nb employé·e·s %d", ent.NbEmployes())
	// Déterminer participants aux formations
	ent.organisationFormation()
	ent.teamBuilding()
	// Envoyer le message aux employés pour qu'ils agissent
	ent.bonneAnnee()
	ent.RecevoirActions(ent.nbAgresseurs + ent.NbEmployes())
	ent.teamBuilding()
	ent.finirCycle()
}

func (ent *Entreprise) stop() {
	ent.fin = true
	go EnvoyerMessageRecrutement(&ent.recrutement, FIN_AGENT, nil)
	for _, emp := range ent.employes {
		go func(emp *Employe) {
			EnvoyerMessage(emp, FIN, nil)
		}(emp)
	}
	ent.RecevoirActions(ent.NbEmployes())
}

func (ent *Entreprise) finirCycle() {
	// // A faire avant GestionDeparts pour bien renvoyer les gens cette année
	ent.gestionPlaintes()

	// Si on le fait en premier, on ne comptera pas ces employés dans les indicateurs ?
	ent.gestionDeparts()
	// A faire en dernier pour ne pas compter les nouveaux employés dans le reste ?
	ent.gestionRecrutements()
	ent.AfficherDonneesCompteur()
	ent.logger.LogType(LOG_ENTREPRISE, "Fin d'année.")
}

// ---------------------
//     Autres
// ---------------------

func (ent *Entreprise) NbEmployes() int {
	return len(ent.employes)
}

func (ent *Entreprise) PourcentageFemmes() float64 {
	femmes := FiltreFemme(ent.employes)
	parite := float64(len(femmes)) / float64(ent.NbEmployes())
	return math.Round(parite*100) / 100
}

func (ent *Entreprise) EnvoyerEmploye(g Genre) *Employe {
	if ent.NbEmployes() == 0 {
		return nil
	}

	var empList []*Employe = nil

	switch g {
	case Homme:
		empList = FiltreHomme(ent.employes)
	case Femme:
		empList = FiltreFemme(ent.employes)
	}

	// si on a personne pour le genre demandé, on va chercher dans l'autre genre
	if len(empList) == 0 {
		idx := rand.Intn(ent.NbEmployes())
		emp := (ent.employes)[idx]
		return emp
	}

	idx := rand.Intn(len(empList))
	emp := empList[idx]
	return emp
}

func (ent *Entreprise) resetCompteur() {
	ent.SetNbEmbauches(0)
	ent.SetNbEmbauchesFemme(0)
	ent.SetNbAgressions(0)
	ent.SetNbPlaintes(0)
	ent.SetNbDemissions(0)
	ent.SetNbDemissionsMaternite(0)
	ent.SetNbRetraites(0)
	ent.SetNbLicenciements(0)
	ent.SetNbDepressions(0)
	ent.SetNbEnfants(0)
	ent.SetNbCongesMaternite(0)
	ent.SetNbCongesPaternite(0)
}

func (ent *Entreprise) AfficherDonneesCompteur() {
	// Mettre des if pour que le log ne s'affiche que si valeur > 0 ? Perte d'info
	ent.logger.LogfType(LOG_RECRUTEMENT, "Recrutement de %d employé·e(s), %d femme(s) et %d homme(s).", ent.NbEmbauches(), ent.NbEmbauchesFemme(), ent.NbEmbauches()-ent.NbEmbauchesFemme())
	if ent.NbAgressions() != 0 {
		ent.logger.LogfType(LOG_AGRESSION, "%d agression(s) sur le lieu de travail dont %d remontée(s) à l'entreprise.", ent.NbAgressions(), ent.NbPlaintes())
	}
	if ent.NbDemissions() != 0 {
		ent.logger.LogfType(LOG_DEPART, "Démission(s) spontanée(s) de %d employé·e(s).", ent.NbDemissions())
	}
	if ent.NbRetraites() != 0 {
		ent.logger.LogfType(LOG_DEPART, "Départ(s) à la retraite pour %d employé·e(s).", ent.NbRetraites())
	}
	if ent.NbPlaintes() != 0 {
		ent.logger.LogfType(LOG_DEPART, "Licenciement pour faute grave appliqué à %d employé·e(s).", ent.NbLicenciements())
	}
	if ent.NbDepressions() != 0 {
		ent.logger.LogfType(LOG_DEPART, "Dépression(s) conduisant à une démission pour %d employé·e(s).", ent.NbDepressions())
	}
	if ent.NbEnfants() != 0 {
		ent.logger.LogfType(LOG_EMPLOYE, "Naissance d'un enfant pour %d employé·e(s).", ent.NbEnfants())
		// S'il n'y a pas d'enfants, il n'y a pas de congés parentaux
		ent.logger.LogfType(LOG_EMPLOYE, "%d employé(s) en congé paternité et %d employée(s) en congé maternité.", ent.NbCongesPaternite(), ent.NbCongesMaternite())
	}
	if ent.NbCongesMaternite() != 0 {
		// S'il n'y a pas de congé maternité, il n'y a pas de démission après congé maternité
		ent.logger.LogfType(LOG_DEPART, "Démission(s) de %d employée(s) après leur congé maternité.", ent.NbDemissionsMaternite())
	}
}

func (ent *Entreprise) MoyenneCompetences() float64 {
	somme := 0
	for _, e := range ent.employes {
		somme += e.Competence()
	}
	return float64(somme) / float64(ent.NbEmployes())
}

func (ent *Entreprise) MoyenneSanteMentale() float64 {
	somme := 0
	for _, e := range ent.employes {
		somme += e.SanteMentale()
	}
	return float64(somme) / float64(ent.NbEmployes())
}
