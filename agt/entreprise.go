package agt

import (
	"log"
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

// ---------------------
//        Général
// ---------------------

func EnvoyerMessageEntreprise(dest *Entreprise, act Action, payload any) {
	dest.chnl <- Communicateur{act, payload}
}

func CollecterActions(dest *Entreprise, act Action, payload any) {
	dest.chnlActions <- Communicateur{act, payload}
}

// ---------------------
//     Constructeur
// ---------------------

// La fonction NewEntreprise doit créer l'entreprise et générer les employés de façon à respecter le quota de parité initial
func NewEntreprise(nbEmployesInit int, pariteInit float64) *Entreprise {
	ent := new(Entreprise)

	var nbFemmes int = int(math.Round(float64(nbEmployesInit) * float64(pariteInit)))
	var nbHommes int = nbEmployesInit - nbFemmes
	var employesInit []Employe

	for i := 0; i < nbFemmes; i++ {
		emp := GenererEmployeInit(&ent, Femme)
		employesInit = append(employesInit, *emp)
		if emp.agresseur {
			ent.nbAgresseurs += 1
		}
	}
	for i := 0; i < nbHommes; i++ {
		emp := GenererEmployeInit(&ent, Homme)
		employesInit = append(employesInit, *emp)
		if emp.agresseur {
			ent.nbAgresseurs += 1
		}
	}
	ent.employes = employesInit
	ent.departs = make([]Employe, 0)
	ent.plaintes = make([][]Employe, 0)
	ent.nbDepressions = 0
	ent.nbRenvois = 0
	ent.ca = 0.0
	ent.nbActions = 0
	ent.chnl = make(chan Communicateur)
	ent.chnlActions = make(chan Communicateur)
	return ent
}

// ---------------------
//        Getters
// ---------------------

func (ent *Entreprise) Employes() []Employe {
	return ent.employes
}

func (ent *Entreprise) Departs() []Employe {
	return ent.departs
}

func (ent *Entreprise) Plaintes() [][]Employe {
	return ent.plaintes
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

	ent.departs = append(ent.departs, *emp)
}

func (ent *Entreprise) RecevoirDepression(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	ent.nbDepressions += 1
	ent.departs = append(ent.departs, *emp)
}

func (ent *Entreprise) RecevoirRetraite(emp *Employe) {
	ent.Lock()
	defer ent.Unlock()

	ent.departs = append(ent.departs, *emp)

	// log.Printf("L'entreprise récupère retraite d'employé %s", emp.Id())
}

func (ent *Entreprise) RecevoirPlainte(plaignant *Employe, accuse *Employe) {
	ent.Lock()
	defer ent.Unlock()

	ent.plaintes = append(ent.plaintes, []Employe{*plaignant, *accuse})

	// log.Printf("L'entreprise récupère plainte d'employé %s contre amployé %s", plaignant.Id(), accuse.Id())
}

// Mettre à jour la formule
func (ent *Entreprise) MettreAJourCA(santeMentale int, competence int) {
	ent.Lock()
	defer ent.Unlock()

	ent.ca += float64(santeMentale) * float64(competence)
}

func (ent *Entreprise) NotifierAction() {
	ent.Lock()
	defer ent.Unlock()

	ent.nbActions += 1

	log.Printf("Nb actions %d/%d", ent.nbActions, (ent.nbEmployes() + ent.nbAgresseurs))

	if ent.nbActions == (ent.nbEmployes() + ent.nbAgresseurs) {
		ent.nbActions = 0
		go CollecterActions(ent, LIBRE, nil)
	}
}

// ---------------------
//     Fin d'année
// ---------------------

// // Renvoyer selon un certain pourcentage
// func (ent Entreprise) gestionPlaintes() {

// }

// func (ent Entreprise) ajusterImpactFemmes() {
// }

// func (ent *Entreprise) calculerBenefice() {
// }

// func (ent *Entreprise) obtenirIndicateursSante() map[string]float64 {
// }

func (ent *Entreprise) gestionDeparts() {
	ent.Lock()
	defer ent.Unlock()

	for _, emp := range ent.departs {
		ent.employes = enleverEmploye(ent.employes, emp)
		if emp.agresseur {
			ent.nbAgresseurs -= 1
		}
	}
}

func (ent *Entreprise) gestionRecrutements() (err error) {
	ent.Lock()
	defer ent.Unlock()

	nbARecruter := float64(ent.nbEmployes()) * constantes.POURCENTAGE_RECRUTEMENT
	embauches, err := ent.recrutement.Recruter(int(math.Round(nbARecruter)))
	if err != nil {
		return err
	}

	for _, emp := range embauches {
		ent.employes = append(ent.employes, emp)
		if emp.agresseur {
			ent.nbAgresseurs += 1
		}
	}
	return nil
}

func (ent *Entreprise) bonneAnnee() {
	ent.Lock()
	defer ent.Unlock()

	ent.nbDepressions = 0
	ent.nbRenvois = 0

	for _, emp := range ent.employes {
		go func(emp Employe) {
			EnvoyerMessage(&emp, LIBRE, nil)
		}(emp)
	}
}

// ---------------------
//  Logique de simulation
// ---------------------

func (ent *Entreprise) Start() {
	for _, emp := range ent.employes {
		go func(emp Employe) {
			emp.Start()
		}(emp)
	}

	// ent.recrutement.Start()

	for {
		ent.agir()
	}
}

func (ent *Entreprise) agir() {
	msg := <-ent.chnl
	if msg.Act == LIBRE {
		log.Printf("Commence l'année")
		// Envoyer le message aux employés pour qu'ils agissent
		ent.bonneAnnee()
		<-ent.chnlActions
		ent.finirCycle()
	}
}

func (ent *Entreprise) finirCycle() {
	// // A faire avant GestionDeparts pour bien renvoyer les gens cette année
	// ent.gestionPlaintes()
	// ent.ajusterImpactFemmes()
	// ent.calculerBenefice()
	// ent.obtenirIndicateursSante()

	// Si on le fait en premier, on ne comptera pas ces employés dans les indicateurs ?
	ent.gestionDeparts()
	// A faire en dernier pour ne pas compter les nouveaux employés dans le reste ?
	ent.gestionRecrutements()
}

// ---------------------
//     Autres
// ---------------------

func (ent *Entreprise) AjouterRecrutement(recrut Recrutement) {
	ent.Lock()
	defer ent.Unlock()

	ent.recrutement = recrut
}

func (ent *Entreprise) nbEmployes() int {
	return len(ent.employes)
}

func (ent *Entreprise) PourcentageFemmes() float64 {
	ent.Lock()
	defer ent.Unlock()

	femmes := FiltreFemme(ent.employes)
	return float64(len(femmes)) / float64(len(ent.employes))
}

func (ent *Entreprise) EnvoyerEmploye(g Genre) *Employe {
	ent.Lock()
	defer ent.Unlock()

	var empList []*Employe = nil

	switch g {
	case Homme:
		empList = FiltreHommePtr(ent.employes)
	case Femme:
		empList = FiltreFemmePtr(ent.employes)
	}

	if empList == nil {
		idx := rand.Intn(len(ent.employes))
		emp := ent.employes[idx]
		return &emp
	}

	idx := rand.Intn(len(empList))
	emp := empList[idx]
	return emp
}
