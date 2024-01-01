// bouton retour fait revenir à l'index
btnRetour.addEventListener("click", () => {
    window.location.href = '../index.html';
});

// Crée un objet qui permet d'afficher des messages
const popupInfo = new InfoPopup();


// Gestion des infos des petits "i"
infosBenefices.addEventListener("mouseover", () => {
    textInfosBenefices.style.display = "block"
});
infosBenefices.addEventListener("mouseleave", () => {
    textInfosBenefices.style.display = "none"
});
infosRecrutement.addEventListener("mouseover", () => {
    textInfosRecrutement.style.display = "block"
});
infosRecrutement.addEventListener("mouseleave", () => {
    textInfosRecrutement.style.display = "none"
});
infosLog.addEventListener("mouseover", () => {
    textInfosLog.style.display = "block"
});
infosLog.addEventListener("mouseleave", () => {
    textInfosLog.style.display = "none"
});



// Initialise le graphe

const TITLE = 0;
const COLOR = 1;

let titles = {
    "Bénéfices" : ["Bénéfices produit par l'entreprise par année", [130, 174, 210]],
    "Parité" : ["Pourcentage de femmes dans l'entreprise", [231, 54, 56]],
    "Compétences" : ["Moyenne des compétences des employé-es (sur 10)", [255, 148, 77]],
    "Santé mentale" : ["Moyenne de la santé mentale des employé-es (sur 100)", [102, 0, 102]]
}

const leGraph = new Graph(document.getElementById('sim-graph'));

leGraph.setIncrement(1);

for(graphName in titles) {
    leGraph.addNewGraph(graphName, titles[graphName][COLOR], titles[graphName][TITLE]);
}

leGraph.selectGraphs("Bénéfices");

btnGraphVisuBenefices.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Bénéfices");
    pressBtn(evt.target);
});

btnGraphVisuParite.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Parité");
    pressBtn(evt.target);
});

btnGraphVisuCompetences.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Compétences");
    pressBtn(evt.target);
});

btnGraphVisuSanteMentale.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Santé mentale");
    pressBtn(evt.target);
});

function pressBtn(btn) {
    for(const domElt of document.getElementsByClassName("btn-visu")) {
        domElt.classList.remove("btn-presse");
    }

    btn.classList.add("btn-presse");
    leGraph.render();
}


// Initialise la taille max de la partie log
document.querySelector(".logs-container").style.maxHeight = document.querySelector(".logs-container").clientHeight + "px";

const leLogger = new Logger(document.getElementById("logger"));

