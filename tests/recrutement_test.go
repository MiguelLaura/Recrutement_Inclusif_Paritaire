package tests

import (
	"testing"

	"gitlab.utc.fr/mennynat/ia04-project/agt"
)

func Test_FiltreFemme(t *testing.T) {
	var ent agt.Entreprise
	var e1 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) //recruté
	var e2 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	var e3 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) //recrutée
	var e4 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	var e5 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	var e6 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 5, &ent) //recrutée
	var employes []agt.Employe
	employes = append(employes, e1)
	employes = append(employes, e2)
	employes = append(employes, e3)
	employes = append(employes, e4)
	employes = append(employes, e5)
	employes = append(employes, e6)
	femmes := agt.FiltreFemme(employes)
	if len(femmes) != 3 {
		t.Errorf("Le slice renvoyé ne contient pas le bon nombre d'élement")
	}
	if !agt.EstDansSliceEmploye(femmes, e3) || !agt.EstDansSliceEmploye(femmes, e5) || !agt.EstDansSliceEmploye(femmes, e6) {
		t.Errorf("Mauvais candidats recrutés: %v", femmes)
	}
}

func Test_FiltreHomme(t *testing.T) {
	var ent agt.Entreprise
	var e1 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) //recruté
	var e2 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	var e3 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) //recrutée
	var e4 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	var e5 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	var e6 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 5, &ent) //recrutée
	var employes []agt.Employe
	employes = append(employes, e1)
	employes = append(employes, e2)
	employes = append(employes, e3)
	employes = append(employes, e4)
	employes = append(employes, e5)
	employes = append(employes, e6)
	hommes := agt.FiltreHomme(employes)
	if len(hommes) != 3 {
		t.Errorf("Le slice renvoyé ne contient pas le bon nombre d'élement")
	}
	if !agt.EstDansSliceEmploye(hommes, e1) || !agt.EstDansSliceEmploye(hommes, e2) || !agt.EstDansSliceEmploye(hommes, e4) {
		t.Errorf("Mauvais candidats recrutés: %v", hommes)
	}
}

func Test_EmployeMaxCompetences(t *testing.T) {
	// TEST 1 : un seul candidat renvoyé
	var ent agt.Entreprise
	var e1 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent)
	var e2 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 4, &ent)
	var e3 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) // renvoyé
	var e4 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 1, &ent)
	var cand []agt.Employe
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)

	embauches := agt.EmployeMaxCompetences(cand)
	if len(embauches) == 0 {
		t.Errorf("Aucun candidat renvoyé")
	}
	if len(embauches) > 1 {
		t.Errorf("Plus d'un candidat renvoyé")
	}
	if embauches[0].Id() != e3.Id() {
		t.Errorf("Le 3e candidat est celui avec la competence max, le candidat retourné est %T", embauches[0])
	}

	// TEST 2: Plusieurs candidats renvoyés

	e1 = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent)
	e2 = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) //renvoyé
	e3 = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) //renvoyé
	e4 = *agt.NewEmploye(1, 0, 100, false, 0.25, 1, &ent)

	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)

	embauches = agt.EmployeMaxCompetences(cand)
	if len(embauches) == 0 {
		t.Errorf("Aucun candidat renvoyé")
	}
	if len(embauches) == 1 {
		t.Errorf("Un seul candidat trouvé: %T", embauches[0])
	}
	if len(embauches) == 2 && ((embauches[0].Id() != e2.Id() && embauches[0].Id() != e3.Id()) || (embauches[1].Id() != e2.Id() && embauches[1].Id() != e3.Id())) {
		t.Errorf("Les mauvais candidats ont été sélectionnés: %T", embauches)
	}
}

func Test_RecrutementCompetencesEgales(t *testing.T) {
	// TEST 1 : Pas d'utilisation des stratégies
	var ent agt.Entreprise
	var e1 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) //recruté
	var e2 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	var e3 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) //recrutée
	var e4 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	var e5 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	var e6 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 5, &ent) //recrutée
	var cand []agt.Employe
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err := agt.RecrutementCompetencesEgales(3, agt.PrioFemme, cand)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 3 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e1) || !agt.EstDansSliceEmploye(embauches, e6) || !agt.EstDansSliceEmploye(embauches, e3) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 2: Utilisation de PrioFemme pour egalité HF
	e1 = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent)
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	e3 = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) // recruté
	e4 = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	e6 = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent) // recrutée
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err = agt.RecrutementCompetencesEgales(2, agt.PrioFemme, cand)

	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 2 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e3) || !agt.EstDansSliceEmploye(embauches, e6) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 3: Utilisation de PrioFemme pour egalité HH

	e1 = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) // 50% de chance d'être recruté
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	e3 = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) // recrutée
	e4 = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	e6 = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) // 50% de chance d'être recruté
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err = agt.RecrutementCompetencesEgales(2, agt.PrioFemme, cand)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 2 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e3) || (!agt.EstDansSliceEmploye(embauches, e6) && !agt.EstDansSliceEmploye(embauches, e1)) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 4: Utilisation de PrioHomme pour egalité HF
	e1 = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) // recruté
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	e3 = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) // recruté
	e4 = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	e6 = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent) // recrutée
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err = agt.RecrutementCompetencesEgales(2, agt.PrioHomme, cand)

	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 2 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e3) || !agt.EstDansSliceEmploye(embauches, e1) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 5: Utilisation de PrioHomme pour egalité FF
	e1 = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent) // 50% de chance d'être recruté
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	e3 = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) // recrutée
	e4 = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	e6 = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent) // 50% de chance d'être recruté
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err = agt.RecrutementCompetencesEgales(2, agt.PrioHomme, cand)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 2 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e3) || (!agt.EstDansSliceEmploye(embauches, e6) && !agt.EstDansSliceEmploye(embauches, e1)) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 6: Erreurs

	cand = nil
	embauches, err = agt.RecrutementCompetencesEgales(2, agt.StratVide, cand)
	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que stratégie inconnue: %s", err)
	}
	embauches, err = agt.RecrutementCompetencesEgales(-1, agt.StratVide, cand)
	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que nombre négatif de personnes à recruter: %s", err)
	}
}

func Test_RecrutementPlacesReservees(t *testing.T) {
	// TEST 1: Nombre de femmes à recruter tombe juste
	var ent agt.Entreprise
	var e1 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) //recruté
	var e2 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent) //recruté
	var e3 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent) //recrutée
	var e4 agt.Employe = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	var e5 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent) //recrutée
	var e6 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 2, &ent)
	var cand []agt.Employe
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err := agt.RecrutementPlacesReservees(4, cand, 0.5)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 4 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e1) || !agt.EstDansSliceEmploye(embauches, e2) || !agt.EstDansSliceEmploye(embauches, e3) || !agt.EstDansSliceEmploye(embauches, e5) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 2: Nombre de femmes à recruter ne tombe pas juste
	e1 = *agt.NewEmploye(0, 0, 100, false, 0.25, 7, &ent) //recruté
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	e3 = *agt.NewEmploye(1, 0, 100, false, 0.25, 9, &ent) //recrutée
	e4 = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent) //recrutée
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 1, &ent)
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)

	embauches, err = agt.RecrutementPlacesReservees(3, cand, 0.5)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 3 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e1) || !agt.EstDansSliceEmploye(embauches, e3) || !agt.EstDansSliceEmploye(embauches, e4) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 3: Il n'y a pas assez de femmes pour atteindre l'objectif donné par le pourcentage
	e1 = *agt.NewEmploye(0, 0, 100, false, 0.25, 8, &ent) //recruté
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent) //recruté
	e3 = *agt.NewEmploye(0, 0, 100, false, 0.25, 5, &ent) //recruté
	e4 = *agt.NewEmploye(0, 0, 100, false, 0.25, 1, &ent)
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 7, &ent) //recrutée
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)

	embauches, err = agt.RecrutementPlacesReservees(4, cand, 0.75)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 4 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e1) || !agt.EstDansSliceEmploye(embauches, e2) || !agt.EstDansSliceEmploye(embauches, e3) || !agt.EstDansSliceEmploye(embauches, e5) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 4: Egalité de compétences
	e1 = *agt.NewEmploye(0, 0, 100, false, 0.25, 6, &ent) //recruté
	e2 = *agt.NewEmploye(0, 0, 100, false, 0.25, 4, &ent)
	e3 = *agt.NewEmploye(0, 0, 100, false, 0.25, 5, &ent) //recruté
	e4 = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent) //recrutée
	e5 = *agt.NewEmploye(1, 0, 100, false, 0.25, 7, &ent) //recrutée
	e6 = *agt.NewEmploye(1, 0, 100, false, 0.25, 3, &ent)
	cand = nil
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	cand = append(cand, e5)
	cand = append(cand, e6)

	embauches, err = agt.RecrutementPlacesReservees(4, cand, 0.25)
	if err != nil {
		t.Errorf("Erreur : %s", err)
	}
	if len(embauches) != 4 {
		t.Errorf("Pas assez de candidats recrutés")
	}
	if !agt.EstDansSliceEmploye(embauches, e1) || !agt.EstDansSliceEmploye(embauches, e3) || !agt.EstDansSliceEmploye(embauches, e4) || !agt.EstDansSliceEmploye(embauches, e5) {
		t.Errorf("Mauvais candidats recrutés: %T", embauches)
	}

	// TEST 5: Gestion des erreurs
	cand = nil
	embauches, err = agt.RecrutementPlacesReservees(4, cand, 12)
	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que pourcentagePlace > 1")
	}
	embauches, err = agt.RecrutementPlacesReservees(4, cand, -0.2)
	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que pourcentagePlace <0")
	}
	embauches, err = agt.RecrutementPlacesReservees(-4, cand, 0.2)
	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que nbARecruter<0")
	}
}

func Test_Recruter(t *testing.T) {
	var ent *agt.Entreprise

	// TEST 1: Erreurs

	r := agt.NewRecrutement(ent, -1, agt.StratVide, agt.StratVide, agt.Competences, agt.Vide, -1, -1)
	_, err := r.Recruter(2)

	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que typeRecrutement=Competences et pas de stratégie de recrutement renseignée")
	}

	r = agt.NewRecrutement(ent, -1, agt.PrioFemme, agt.PrioHomme, agt.Competences, agt.Vide, -1, -1)
	_, err = r.Recruter(2)

	if err == nil {
		t.Errorf("Pas d'erreur renvoyée alors que pas d'objectif mais deux stratégies renseignées ")
	}

}
