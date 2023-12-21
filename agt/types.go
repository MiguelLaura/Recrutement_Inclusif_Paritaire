package agt

import (
	"sync"
	"time"

	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

// ------------ Types de logs --------------

const (
	LOG_AGRESSION   logger.LogType = "agression"
	LOG_DEPART      logger.LogType = "depart"
	LOG_RECRUTEMENT logger.LogType = "recrutement"
	LOG_EMPLOYE     logger.LogType = "employe"
	LOG_ENTREPRISE  logger.LogType = "entreprise"
	LOG_EVENEMENT   logger.LogType = "evenement"
	LOG_GLOBALE     logger.LogType = "globale"
	LOG_REPONSE     logger.LogType = "reponse"
)

// ------------ SIMULATION ------------

type EtatSimulation struct {
	nbEmp  int
	parite float64
}

type SimulationLocker struct {
	sync.WaitGroup
	sync.Mutex
}

type Simulation struct {
	ent        Entreprise
	pariteInit float64
	maxStep    int
	step       int // Stats
	start      time.Time
	status     Status // created, started, pause, finished
	logger     logger.Loggers
	etatInit   EtatSimulation
	locker     SimulationLocker
	aCommencee bool
}

type Action int

// Action est une enumeration
const (
	NOOP Action = iota
	LIBRE
	AGRESSION
	FORMATION
	FIN
)

type Status int

// Action est une enumeration
const (
	CREATED Status = iota
	STARTED
	PAUSED
	ENDED
	STEP
)

// Permet la communication entre entreprise et employé
type Communicateur struct {
	Act     Action
	Payload any
}

type Action_recrutement int

const (
	RECRUTEMENT Action_recrutement = iota
	FIN_RECRUTEMENT
	ERREUR_RECRUTEMENT
	FIN_AGENT
)

// Permet la communication entre agents
type Communicateur_recrutement struct {
	Act     Action_recrutement
	Payload any
}

// ------------ EMPLOYE ------------

type EmployeID string

var agtCnt int = 0

type Employe struct {
	id              EmployeID
	genre           Genre
	anciennete      int //entre 0 et 40
	santeMentale    int //entre 0 et 100
	agresseur       bool
	competence      int //entre 0 et 10
	cmpt_competence int // entre 0 et 5. Quand il atteint 5, competence +1
	entreprise      *Entreprise
	fin             bool
	chnl            chan Communicateur
	logger          *logger.Loggers
}

type Genre int

const (
	Homme Genre = iota
	Femme
)

// ------------ ENTREPRISE ------------

type Entreprise struct {
	sync.Mutex
	employes        *[]Employe
	departs         *[]Employe
	formation       *[]Employe
	conge_parental  *[]Employe
	plaintes        *[][]Employe
	nbDepressions   int
	nbRenvois       int
	recrutement     Recrutement
	ca              float64
	nbActions       int
	nbAgresseurs    int
	fin             bool
	chnl            chan Communicateur
	chnlActions     chan Communicateur
	chnlRecrutement chan Communicateur_recrutement
	chnlNotifAction chan Communicateur
	logger          *logger.Loggers
}

// ------------ RECRUTEMENT ------------

type StratParite int

const (
	StratVide StratParite = iota // décrit l'absence de StartParite = 0
	PrioHomme                    // = 1
	PrioFemme                    // = 2
	Hasard                       // = 3
)

type TypeRecrutement int

const (
	Vide            TypeRecrutement = iota // décrit l'absence de TypeRecrutement = 0
	Competences                            // = 1
	PlacesReservees                        // = 2
)

type Recrutement struct {
	entreprise             *Entreprise
	objectif               float64     // -1 si non renseigné, entre 0 et 1 sinon
	stratAvant             StratParite // stratVide si non renseigné
	stratApres             StratParite
	typeRecrutementAvant   TypeRecrutement // Vide si non renseigné
	typeRecrutementApres   TypeRecrutement
	pourcentagePlacesAvant float64 // -1 si non renseigné, entre 0 et 1 sinon
	pourcentagePlacesApres float64
	chnl                   chan Communicateur_recrutement
	fin                    bool
	logger                 *logger.Loggers
}
