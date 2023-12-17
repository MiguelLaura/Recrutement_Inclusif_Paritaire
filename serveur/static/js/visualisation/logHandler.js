const logContainer = document.getElementById("logger");

function addLog(contenu) {
    const newLog = document.createElement("p");
    newLog.classList.add("log");

    const now = new Date();
    newLog.textContent = `${now.toTimeString().split(" ")[0]} : ${contenu}`;

    logContainer.appendChild(newLog);
}

