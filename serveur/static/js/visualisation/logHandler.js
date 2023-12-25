const logContainer = document.getElementById("logger");

function addLog(type, contenu) {
    const newLog = document.createElement("p");
    newLog.classList.add("log");
    newLog.classList.add(type);
    let moveScroll = logContainer.scrollTop == logContainer.scrollHeight - logContainer.clientHeight;

    const now = new Date();
    newLog.textContent = `${now.toTimeString().split(" ")[0]} : ${contenu}`;

    logContainer.appendChild(newLog);

    // Si on est en bas de la barre de défilement, on l'ajuste au nouveau log ajouté
    if(moveScroll) {
        logContainer.scrollTop = logContainer.scrollHeight - logContainer.clientHeight;
    }
}

function resetLogs() {
    logContainer.innerHTML = "";
}

/*

    Le menu des options

*/

containerOpt.addEventListener("mouseover", () => {
    if(containerOpt.dataset.toggle === "f") {
        containerOpt.dataset.toggle = "t";
        btnOpt.classList.remove("bi-chevron-down");
        btnOpt.classList.add("bi-chevron-up");
    }
});

containerOpt.addEventListener("mouseleave", () => {
    if(containerOpt.dataset.toggle === "t") {
        containerOpt.dataset.toggle = "f";
        btnOpt.classList.add("bi-chevron-down");
        btnOpt.classList.remove("bi-chevron-up");
    }
});
