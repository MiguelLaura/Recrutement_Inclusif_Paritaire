// Initialise le graphe

const leGraph = new Graph(document.getElementById('sim-graph'));

leGraph.setIncrement(1);
leGraph.addNewGraph("Bénéfices", [130, 174, 210]);
leGraph.addNewGraph("Parité", [231, 54, 56]);

leGraph.selectGraphs("Bénéfices", "Parité");

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

const leLogger = new Logger(document.getElementById("logger"));

