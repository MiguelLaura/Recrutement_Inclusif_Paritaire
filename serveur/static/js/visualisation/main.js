// Ajout des logs

const exLogs = [
    "Employé #7 (H) a agressé Employé #0777 (F)",
    "Employé #0777 (F) a porté plainte contre Employé #7 (H)",
    "Employé #76 (F) a posé sa démission",
]

for(let log of exLogs) {
    addLog(log);
}

// Initialise le graphe

const leGraph = new Graph(document.getElementById('sim-graph'));

leGraph.setIncrement(1);
leGraph.addNewGraph("Bénéfices", [130, 174, 210]);
leGraph.addNewGraph("Parité", [231, 54, 56]);

leGraph.selectGraphs("Bénéfices", "Parité");

// const dataBenef = [-2500, -1900, -10000, -4950, -4950, -3970, -5300];
// const dataParite = [0, 600, 600, 1000, 500, 500, 1300];

// for(let idxData = 0; idxData < dataBenef.length; idxData++) {
//     leGraph.addData(dataBenef[idxData], dataParite[idxData]);
// }

btnGraphVisuTout.addEventListener("click", (evt) => {
    leGraph.selectGraphs("Bénéfices", "Parité");
    pressBtn(evt.target);
});

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
