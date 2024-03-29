// Les champs avec les informations initiales
const anneesInit = document.getElementById("annees-init");
const objectif = document.getElementById("objectif");
const recrutement = document.getElementById("recrutement");
const nbEmployesInit = document.getElementById("nb-employes-init");
const pariteInit = document.getElementById("parite-init");

// Les champs à actualiser à chaque nouvelle année
const anneeElt = document.getElementById("annee");
const nbEmpElt = document.getElementById("nb-emp");
const pariteElt = document.getElementById("parite"); //parité actuelle
const benefice = document.getElementById("benefice") //bénéfice indicateur

// Les boutons pour selectionner le graphe que l'on souhaite
const btnGraphVisuTout = document.getElementById("visu-tout");
const btnGraphVisuBenefices = document.getElementById("visu-benef");
const btnGraphVisuParite = document.getElementById("visu-parite");
const btnGraphVisuCompetences = document.getElementById("visu-competences");
const btnGraphVisuSanteMentale = document.getElementById("visu-sante-mentale");

// Les informations sur la simulation
const statusSimu = document.getElementById("status-simu");
const idNumberSimu = document.getElementById("id-number-simu");

// Bouton retour vers le formulaire
const btnRetour = document.getElementById("button-retour");

// Les options de log
const menuOpt = document.getElementById("dropdown-menu");
const btnOpt = document.querySelector(".dropdown-icon");
const containerOpt = document.querySelector(".dropdown-container");

const checkOpts = document.querySelectorAll(".dropdown-menu input");

// Les informations en plus (dans les petits i)
const infosBenefices = document.getElementById("infos-benefice")
const textInfosBenefices = document.getElementById("text-infos-benefice")
const infosRecrutement = document.getElementById("infos-recrutement")
const textInfosRecrutement = document.getElementById("text-infos-recrutement")
const infosLog = document.getElementById("infos-log")
const textInfosLog = document.getElementById("text-infos-log")