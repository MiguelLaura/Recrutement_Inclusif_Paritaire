class Logger {

    static HISTORY_SIZE = 10000;

    constructor(parent) {
        this.hiddenCategories = new Set();
        this.parent = parent;
    }

    addLog(type, contenu, logTime = (new Date()).toTimeString().split(" ")[0]) {
        const newLog = document.createElement("p");
        newLog.classList.add("log");
        newLog.classList.add(type);
        if(this.hiddenCategories.has(type)) {
            newLog.classList.add("hidden");
        }
        newLog.dataset.logCat = type;
        let moveScroll = this.parent.scrollTop == this.parent.scrollHeight - this.parent.clientHeight;
        
        newLog.textContent = `${logTime} : ${contenu}`;
        
        const otherLogs = this.parent.getElementsByTagName("p");
        if(otherLogs.length >= Logger.HISTORY_SIZE) {
            otherLogs[0].remove();
        }

        this.parent.appendChild(newLog);
    
        // Si on est en bas de la barre de défilement, on l'ajuste au nouveau log ajouté
        if(moveScroll) {
            this.parent.scrollTop = this.parent.scrollHeight - this.parent.clientHeight;
        }
    }

    showLogs(...types) {
        types.forEach(t => this.hiddenCategories.delete(t));
        for(const elt of this.parent.getElementsByTagName("p")) {
            if(types.indexOf(elt.dataset.logCat) !== -1)
                elt.classList.remove("hidden");
        }
    }

    hideLogs(...types) {
        types.forEach(t => this.hiddenCategories.add(t));
        for(const elt of this.parent.getElementsByTagName("p")) {
            if(types.indexOf(elt.dataset.logCat) !== -1)
                elt.classList.add("hidden");
        }
    }

    reset() {
        this.parent.innerHTML = "";
    }
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

checkOpts.forEach(checkElt => {
    checkElt.addEventListener("change", (evt) => {
        const checkElt = evt.target;
        const logType = checkElt.dataset.logtype;

        console.log(checkElt.checked ? "show" : "hide", logType);
        
        if(checkElt.checked) {
            leLogger.showLogs(logType);
        } else {
            leLogger.hideLogs(logType);
        }
    });
});