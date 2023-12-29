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

const leGraph = new Graph(document.getElementById('sim-graph'));

leGraph.setIncrement(1);
leGraph.addNewGraph("Bénéfices", [130, 174, 210]);
leGraph.addNewGraph("Parité", [231, 54, 56]);

leGraph.selectGraphs("Bénéfices");

btnGraphVisuBenefices.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Bénéfices");
    pressBtn(evt.target);
});

btnGraphVisuParite.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Parité");
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

