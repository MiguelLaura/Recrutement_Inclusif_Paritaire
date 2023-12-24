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