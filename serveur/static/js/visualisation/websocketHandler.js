NOT_STARTED = "0";
PLAYING = "1";
PAUSED = "2";

// Crée un objet qui permet d'afficher des messages

const popupInfo = new InfoPopup();

// -------------------------
// WebSocket

const btnToggle = document.getElementById("toggle-simu");
const btnStop = document.getElementById("stop-simu");
const btnRelancer = document.getElementById("restart-simu");

let conn = undefined;

const currentURL = window.location.href;
const url = new URL(currentURL);
const pathname = url.pathname;
const pathParts = pathname.split('/');
const id = pathParts[pathParts.length - 1];

const numbersId = id.split("_")
const numberId = numbersId[numbersId.length - 1]
idNumberSimu.innerText = numberId; //mettre le numéro de la simulation dans le titre



btnToggle.addEventListener("click", () => {
    switch(btnToggle.dataset.state) {
        case NOT_STARTED: // Pas encore démarrée
            conn.send(JSON.stringify({id_simulation : id, type: "action", data: "start"}));
        break;
        case PLAYING: // En train de simuler
            conn.send(JSON.stringify({id_simulation : id, type: "action", data: "pause"}));
        break;
        case PAUSED: // En pause
            conn.send(JSON.stringify({id_simulation : id, type: "action", data: "continue"}));
    }
});

btnStop.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "stop"}));
});

btnRelancer.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "relancer"}));
});

if (window["WebSocket"]) {
    console.log("supports websockets");
    // Connect to websocket
    conn = new WebSocket("ws://" + document.location.host + "/ws");
} else {
    popupInfo.error("Not supporting websockets", 5);
}

conn.addEventListener("open", () => {
    popupInfo.info("connected", 1);
    //pour dire que l'on souhaite récupérer les informations sur la simulation
    conn.send(JSON.stringify({id_simulation : id, type: "init", data: ""}));
});

conn.addEventListener("close", () => {
    console.log("disconnected");
});

conn.addEventListener("message", (evt) => {
    resp = JSON.parse(evt.data);
    
    if("data" in resp) {
        switch(resp.type) {
            case "erreur":
                popupInfo.error(resp.data, 10);
                break;
            case "info":
                popupInfo.info(resp.data, 10);
                break;
            case "globale":
                const data = resp.data[0];
                mettreLesChosesAuBonEndroit(data);
                break;
            case "initial":
                traiterReponseAction(resp.data[0].status, true) //utiliser le status de la simulation
                afficherInformationsInitiales(resp.data[0])
                break;
            case "reponse": 
                traiterReponseAction(resp.data[0].action, resp.data[0].succes)
                break;
            default:
                addLog(`[${resp.type}] ${resp.data}`)
        }
    }
});

//fonction qui met à jour les informations de la simulation en cours du temps
function mettreLesChosesAuBonEndroit(data) {
    anneeElt.textContent = data.annee + " ans";
    nbEmpElt.textContent = data.nbEmp + " employé·es";
    pariteElt.textContent = (data.parite * 100).toFixed(0) + "% de femmes";
    benefice.textContent = data.benefices + "€"

    leGraph.addData(data.benefices.toFixed(0), (data.parite * 100).toFixed(0));
}

function afficherInformationsInitiales(data) {
    nbEmployesInit.innerText = "(début : "+data.nbEmployesInit+")"
    pariteInit.innerText = "(début : "+(data.pariteInit * 100).toFixed(0)+"%)"

    if (data.objectif != "-1") {
        objectif.innerText = "Objectif : "+(data.objectif * 100).toFixed(0)+"%"
    } //sinon, pas de texte pour l'objectif

    textHTMLrecrutement = "<li>"+data.recrutAvant+"</li>"
    if(data.recrutApres != "") { //s'il y a un recrutement avant
        textHTMLrecrutement += "<li>"+data.recrutApres+"</li>"
    }
    recrutement.innerHTML = textHTMLrecrutement
}

function resetData() {
    const dataElements = document.querySelectorAll('.data');
    dataElements.forEach(span => span.innerHTML = "");
}

function traiterReponseAction(action, succes) {
    switch(action) {
        case "start":
            if(succes) {
                btnToggle.firstChild.classList.toggle("bi-play-fill");
                btnToggle.firstChild.classList.toggle("bi-pause-fill");
                btnToggle.lastChild.textContent = "Pause";
                btnToggle.dataset.state = PLAYING;
                statusSimu.innerText = "[en cours]";
            }
            break;
        case "pause":
            if(succes) {
                btnToggle.firstChild.classList.toggle("bi-play-fill");
                btnToggle.firstChild.classList.toggle("bi-pause-fill");
                btnToggle.lastChild.textContent = "Reprendre";
                btnToggle.dataset.state = PAUSED;
                statusSimu.innerText = "[en pause]";
            }
            break;
        case "continue":
            if(succes) {
                btnToggle.firstChild.classList.toggle("bi-play-fill");
                btnToggle.firstChild.classList.toggle("bi-pause-fill");
                btnToggle.lastChild.textContent = "Pause";
                btnToggle.dataset.state = PLAYING;
                statusSimu.innerText = "[en cours]";
            }
            break;
        case "stop":
            if(succes) {
                btnToggle.firstChild.classList.add("bi-play-fill");
                statusSimu.innerText = "[terminée]";
            }
            break;
        case "relancer":
            if(succes) {
                btnToggle.firstChild.classList.add("bi-play-fill");
                btnToggle.lastChild.textContent = "Commencer";
                btnToggle.dataset.state = NOT_STARTED;
                statusSimu.innerText = "[pas débutée]";
                resetLogs();
                resetData();

                // TODO : A voir si on reset le graph une fois la simulation relancée
                leGraph.reset();
                leGraph.render();
            }
            break;
        case "not_started" : 
            if(succes) {
                btnToggle.firstChild.classList.add("bi-play-fill");
                btnToggle.lastChild.textContent = "Commencer";
                btnToggle.dataset.state = NOT_STARTED;
                statusSimu.innerText = "[pas débutée]";
            }
        break;
    }
}