package agt

import (
	"errors"
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

func CollecterActions(dest *Entreprise, act Action, payload any) {
	dest.chnlActions <- Communicateur{act, payload}
}

func EnvoyerMessageRecrutement(dest *Recrutement, act Action_recrutement, payload any) {
	dest.chnl <- Communicateur_recrutement{act, payload}
}

// ---------------------
//     Constructeur
// ---------------------

// La fonction NewEntreprise doit créer l'entreprise et générer les employés de façon à respecter le quota de parité initial
func NewEntreprise(nbEmployesInit int, pariteInit float64, logger *logger.Loggers) *Entreprise {
	ent := new(Entreprise)

	var nbFemmes int = int(math.Round(float64(nbEmployesInit) * float64(pariteInit)))
	var nbHommes int = nbEmployesInit - nbFemmes
	var employesInit []Employe

	for i := 0; i < nbFemmes; i++ {
		emp := GenererEmployeInit(&ent, Femme, logger)
		employesInit = append(employesInit, *emp)
		if emp.agresseur {
			ent.nbAgresseurs += 1
		}
	}
	for i := 0; i < nbHommes; i++ {
		emp := GenererEmployeInit(&ent, Homme, logger)
		employesInit = append(employesInit, *emp)
		if emp.agresseur {
			ent.nbAgresseurs += 1
		}
	}
	ent.employes = &employesInit
	departs := make([]Employe, 0)
	ent.departs = &departs
	formation := make([]Employe, 0)
	ent.formation = &formation
	conge_parental := make([]Employe, 0)
	ent.conge_parental = &conge_parental
	plaintes := make([][]Employe, 0)
	ent.plaintes = &plaintes
	ent.nbDepressions = 0
	ent.nbRenvois = 0
	ent.ca = 0.0
	ent.nbActions = 0
	ent.fin = false
	ent.chnl = make(chan Communicateur)
	ent.chnlActions = make(chan Communicateur)
	ent.chnlRecrutement = make(chan Communicateur_recrutement)
	ent.chnlNotifAction = make(chan Communicateur)
	ent.logger = logger
	return ent
}

// ---------------------
//        Getters
// ---------------------

func (ent *Entreprise) Employes() []Employe {
	return *ent.employes
}

func (ent *Entreprise) Departs() []Employe {
	return *ent.departs
}

func (ent *Entreprise) Formation() []Employe {
	return *ent.formation
}

func (ent *Entreprise) Plaintes() [][]Employe {
	return *ent.plaintes
}

func (ent *Entreprise) NbDepressions() int {
	return ent.nbDepressions
}

func (ent *Entreprise) NbRenvois() int {
	return ent.nbRenvois
}

func (ent *Entreprise) Recrutement() Recrutement {
	return ent.recrutement
}

func (ent *Entreprise) Ca() float64 {
	return ent.ca
}

func (ent *Entreprise) NbActions() int {
	return ent.nbActions
}

func (ent *Entreprise) NbAgresseurs() int {
	return ent.nbAgresseurs
}

// ---------------------
//     En cours d'année (appelées par les employés)
// ---------------------

func (ent *Entreprise) RecevoirDemission(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	i, _ := TrouverEmploye(*ent.departs, func(e Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		*ent.departs = append(*ent.departs, *emp)
		ent.logger.LogfType(LOG_DEPART, "Demission %s : nb départs %d", emp.Id(), len(*ent.departs))
		return
	}
}

func (ent *Entreprise) RecevoirDepression(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	ent.nbDepressions += 1
	i, _ := TrouverEmploye(*ent.departs, func(e Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		*ent.departs = append(*ent.departs, *emp)
		ent.logger.LogfType(LOG_DEPART, "Depression %s : nb départs %d", emp.Id(), len(*ent.departs))
		return
	}
}

func (ent *Entreprise) RecevoirRetraite(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	i, _ := TrouverEmploye(*ent.departs, func(e Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		*ent.departs = append(*ent.departs, *emp)
		ent.logger.LogfType(LOG_DEPART, "Retraite %s : nb départs %d", emp.Id(), len(*ent.departs))
		return
	}
}

func (ent *Entreprise) CongeParental(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()
	i, _ := TrouverEmploye(*ent.conge_parental, func(e Employe) bool { return e.Id() == emp.Id() }, 0)
	if i < 0 {
		*ent.conge_parental = append(*ent.conge_parental, *emp)
		ent.logger.LogfType(LOG_ENTREPRISE, "Conge parental %s", emp.Id())
		return
	}
}

func (ent *Entreprise) RecevoirPlainte(plaignant *Employe, accuse *Employe) {
	ent.Lock()
	defer ent.Unlock()

	*ent.plaintes = append(*ent.plaintes, []Employe{*plaignant, *accuse})
}

// Mettre à jour la formule
func (ent *Entreprise) MettreAJourCA(santeMentale int, competence int) {
	ent.Lock()
	defer ent.Unlock()

	ent.ca += float64(santeMentale) * float64(competence)
}

func (ent *Entreprise) RecevoirActions(nbActions int) {
	for {
		<-ent.chnlNotifAction

		ent.nbActions += 1

		// ent.logger.LogfType(LOG_ENTREPRISE, "Nb actions %d/%d", ent.nbActions, nbActions)

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
	ent.logger.LogType(LOG_ENTREPRISE, "Organisation team building")
	for _, e := range *ent.employes {
		test_presence, _ := TrouverEmploye(*ent.departs, func(emp Employe) bool { return e.Id() == emp.Id() }, 0)
		// On vérifie que l'employé ne va pas partir
		if test_presence < 0 {
			if e.santeMentale < 100 {
				if e.santeMentale+constantes.BOOST_TEAM_BUILDING > 100 {
					e.santeMentale = 100
				} else {
					e.santeMentale += constantes.BOOST_TEAM_BUILDING
				}
			}
		}
	}
}

func (ent *Entreprise) organisationFormation() {
	ent.logger.LogType(LOG_ENTREPRISE, "Organisation formation")

	*ent.formation = make([]Employe, 0)
	// Génération des employés participant à une formation cette année

	// 32% des français ont participé à une formation
	nb_employes_formes := math.Round(constantes.POURCENTAGE_FORMATION * float64(ent.nbEmployes()))
	// 50% des employés qui se forment sont des femmes
	nb_femmes_formes := math.Round(nb_employes_formes / 2)
	nb_hommes_formes := nb_femmes_formes
	femmes := FiltreFemme(*ent.employes)
	hommes := FiltreHomme(*ent.employes)
	for idx := 0; idx < int(nb_femmes_formes); idx++ {
		if len(femmes) == 0 {
			break
		}
		i := rand.Intn(len(femmes))
		*ent.formation = append(*ent.formation, femmes[i])
		// Pour ne pas avoir de doublons
		femmes = enleverEmploye(femmes, femmes[i])
	}
	for idx := 0; idx < int(nb_hommes_formes); idx++ {
		if len(hommes) == 0 {
			break
		}
		i := rand.Intn(len(hommes))
		*ent.formation = append(*ent.formation, hommes[i])
		// Pour ne pas avoir de doublons
		hommes = enleverEmploye(hommes, hommes[i])
	}
	ent.logger.LogfType(LOG_ENTREPRISE, "nb employe: %d nb_formes: %d", ent.nbEmployes(), len(*ent.formation))
}

// ---------------------
//     Fin d'année
// ---------------------

// // Renvoyer selon un certain pourcentage
func (ent *Entreprise) gestionPlaintes() {
	if len(*ent.plaintes) <= 0 {
		return
	}
	for _, e := range *ent.plaintes {
		if rand.Float64() <= constantes.PROBA_LICENCIEMENT {
			accuse := e[1]
			i, _ := TrouverEmploye(*ent.departs, func(e Employe) bool { return e.Id() == accuse.Id() }, 0)
			if i < 0 {
				*ent.departs = append(*ent.departs, accuse)
				ent.logger.LogfType(LOG_DEPART, "Licenciement %s : nb départs %d", accuse.Id(), len(*ent.departs))
			}
		}
	}
	*ent.plaintes = make([][]Employe, 0)
}

func (ent *Entreprise) calculerBenefice() (benef float64) {
	benef = 0

	// Pour chaque employé, on calcule ce qu'il rapporte à l'entreprise en fonction de sa santé mentale et compétences
	// Compétences varient entre 0 et 10 par une loi normale d'espérance 5.
	// CA_PAR_EMPLOYE/5 * competences pour garder la valeur moyenne du CA_PAR_EMPLOYE mais prendre en compte les compétences
	// benef plus faible si santé mentale plus basse

	for _, e := range *ent.employes {
		benef += (constantes.CA_PAR_EMPLOYE/5)*float64(e.competence)*float64(e.santeMentale)/100 - constantes.COUT_EMPLOYE
	}

	// Impact conges parental : pas de salaire et pas de productivité pendant une certaine période
	for _, e := range *ent.conge_parental {
		if e.Genre() == Femme {
			benef -= constantes.PROPORTION_ARRET_F * ((constantes.CA_PAR_EMPLOYE/5)*float64(e.competence)*float64(e.santeMentale)/100 - constantes.COUT_EMPLOYE)
		} else {
			benef -= constantes.PROPORTION_ARRET_H * ((constantes.CA_PAR_EMPLOYE/5)*float64(e.competence)*float64(e.santeMentale)/100 - constantes.COUT_EMPLOYE)
		}
	}

	// Bonus de productivité si %femmes supérieur à 35%
	if ent.PourcentageFemmes() > constantes.SEUIL_IMPACT_FEMME {
		benef = benef * (1.0 + constantes.BOOST_PRODUCTIVITE_FEMME)
	}

	// Coût du recrutement
	nbARecruter := float64(ent.nbEmployes())*constantes.POURCENTAGE_RECRUTEMENT + 1.0
	benef -= float64(nbARecruter * constantes.COUT_RECRUTEMENT)

	// Coût des formations
	nbFormes := len(*ent.formation)
	benef -= float64(constantes.PRIX_FORMATION * constantes.NB_JOURS_FORMATION * nbFormes)

	// Amende si non parité
	// Modèle le plus simple : si %Femmes ne respectent pas la loi (<40%), amende d'1% des bénéfices
	// Modèle 2 plus proche de la réalité : amende si non respect pendant 3 ans consécutifs
	// Modèle 3 le plus réaliste : amende à partir de 2029
	if ent.PourcentageFemmes() < constantes.SEUIL_AMENDE {
		benef = benef * (1 - constantes.POURCENTAGE_AMENDE)
	}

	return benef
}

func (ent *Entreprise) gestionDeparts() {
	if len(*ent.departs) <= 0 {
		return
	}
	for _, emp := range *ent.departs {
		*ent.employes = enleverEmploye(*ent.employes, emp)
		go func(emp Employe) {
			EnvoyerMessage(&emp, FIN, nil)
		}(emp)
		if emp.agresseur {
			ent.nbAgresseurs -= 1
		}
	}
	ent.RecevoirActions(len(*ent.departs))
	*ent.departs = make([]Employe, 0)
}

func (ent *Entreprise) gestionRecrutements() (err error) {
	msg := <-ent.chnlRecrutement
	if msg.Act == ERREUR_RECRUTEMENT {
		return msg.Payload.(error)
	} else if msg.Act == FIN_RECRUTEMENT {
		embauches := msg.Payload.([]Employe)
		ent.logger.LogfType(LOG_ENTREPRISE, "Embauche %d, employés %d", len(embauches), ent.nbEmployes())
		for _, emp := range embauches {
			*ent.employes = append(*ent.employes, emp)
			if emp.agresseur {
				ent.nbAgresseurs += 1
			}
		}
		return nil
	}

	err = errors.New("erreur : erreur recrutement")
	return err
}

func (ent *Entreprise) lancerRecrutements() {
	nbARecruter := float64(ent.nbEmployes())*constantes.POURCENTAGE_RECRUTEMENT + 1.0

	go EnvoyerMessageRecrutement(&ent.recrutement, RECRUTEMENT, int(nbARecruter))
}

func (ent *Entreprise) bonneAnnee() {
	ent.nbDepressions = 0
	ent.nbRenvois = 0
	*ent.conge_parental = make([]Employe, 0)

	for _, emp := range *ent.employes {
		go func(emp Employe) {
			EnvoyerMessage(&emp, LIBRE, nil)
		}(emp)
	}

	ent.lancerRecrutements()
}

// ---------------------
//  Logique de simulation
// ---------------------

func (ent *Entreprise) Start() {
	for _, emp := range *ent.employes {
		go func(emp Employe) {
			emp.Start()
		}(emp)
	}

	go ent.recrutement.Start()

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
	ent.logger.LogType(LOG_ENTREPRISE, "Fin d'entreprise")
	<-ent.chnl
}

func (ent *Entreprise) agir() {
	if len(*ent.employes) <= 0 {
		ent.fin = true
		return
	}
	ent.logger.LogType(LOG_ENTREPRISE, "Commence l'année")
	ent.logger.LogfType(LOG_ENTREPRISE, "Nb employe %d", ent.nbEmployes())
	// Déterminer participants aux formations
	ent.organisationFormation()
	ent.teamBuilding()
	// Envoyer le message aux employés pour qu'ils agissent
	ent.bonneAnnee()
	ent.RecevoirActions(ent.nbAgresseurs + ent.nbEmployes())
	ent.teamBuilding()
	ent.finirCycle()
}

func (ent *Entreprise) stop() {
	ent.fin = true
	go EnvoyerMessageRecrutement(&ent.recrutement, FIN_AGENT, nil)
	for _, emp := range *ent.employes {
		go func(emp Employe) {
			EnvoyerMessage(&emp, FIN, nil)
		}(emp)
	}
	ent.RecevoirActions(ent.nbEmployes())
}

func (ent *Entreprise) finirCycle() {
	// // A faire avant GestionDeparts pour bien renvoyer les gens cette année
	ent.gestionPlaintes()
	benef := ent.calculerBenefice()
	ent.logger.LogfType(LOG_ENTREPRISE, "Benefices: %.2f", benef)
	// ent.obtenirIndicateursSante()
	moy := ent.MoyenneCompetences()
	ent.logger.LogfType(LOG_ENTREPRISE, "moyenne competences %.2f", moy)
	// Si on le fait en premier, on ne comptera pas ces employés dans les indicateurs ?
	ent.gestionDeparts()
	// A faire en dernier pour ne pas compter les nouveaux employés dans le reste ?
	ent.gestionRecrutements()
	ent.logger.LogType(LOG_GLOBALE, SituationActuelle{ent.nbEmployes(), ent.PourcentageFemmes(), benef})
	ent.logger.LogType(LOG_ENTREPRISE, "Fin d'année\n\n")

}

// ---------------------
//     Autres
// ---------------------

func (ent *Entreprise) AjouterRecrutement(recrut Recrutement) {
	ent.recrutement = recrut
}

func (ent *Entreprise) nbEmployes() int {
	return len(*ent.employes)
}

func (ent *Entreprise) PourcentageFemmes() float64 {
	femmes := FiltreFemme(*ent.employes)
	return float64(len(femmes)) / float64(len(*ent.employes))
}

func (ent *Entreprise) EnvoyerEmploye(g Genre) *Employe {
	if len(*ent.employes) == 0 {
		return nil
	}

	var empList []*Employe = nil

	switch g {
	case Homme:
		empList = FiltreHommePtr(*ent.employes)
	case Femme:
		empList = FiltreFemmePtr(*ent.employes)
	}

	// si on a personne pour le genre demandé, on va chercher dans l'autre genre
	if len(empList) == 0 {
		idx := rand.Intn(len(*ent.employes))
		emp := (*ent.employes)[idx]
		return &emp
	}

	idx := rand.Intn(len(empList))
	emp := empList[idx]
	return emp
}

func (ent *Entreprise) MoyenneCompetences() float64 {
	somme := 0
	for _, e := range *ent.employes {
		somme += e.competence
	}
	return float64(somme / ent.nbEmployes())
}
