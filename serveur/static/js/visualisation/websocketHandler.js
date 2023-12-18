NOT_STARTED = "0";
PLAYING = "1";
PAUSED = "2";

// Crée un objet qui permet d'afficher des messages

const popupInfo = new InfoPopup();

// -------------------------
// WebSocket

const btnToggle = document.getElementById("toggle-simu");
const btnStop = document.getElementById("stop-simu");

let conn = undefined;

const currentURL = window.location.href;
const url = new URL(currentURL);
const pathname = url.pathname;
const pathParts = pathname.split('/');
const id = pathParts[pathParts.length - 1];
console.log("ID extrait de l'URL :", id);


//ENVOYER ICI EN PRENANT window.location.href ????

btnToggle.addEventListener("click", () => {
    switch(btnToggle.dataset.state) {
        case NOT_STARTED: // Pas encore démarrée
            conn.send(JSON.stringify({id_simulation : id, type: "action", data: "start"}));
            btnToggle.lastChild.textContent = "Pause";
            btnToggle.dataset.state = PLAYING;
        break;
        case PLAYING: // En train de simuler
            conn.send(JSON.stringify({id_simulation : id, type: "action", data: "pause"}));
            btnToggle.lastChild.textContent = "Reprendre";
            btnToggle.dataset.state = PAUSED;
        break;
        case PAUSED: // En pause
            conn.send(JSON.stringify({id_simulation : id, type: "action", data: "continue"}));
            btnToggle.lastChild.textContent = "Pause";
            btnToggle.dataset.state = PLAYING;
    }

    btnToggle.firstChild.classList.toggle("bi-play-fill");
    btnToggle.firstChild.classList.toggle("bi-pause-fill");
});

btnStop.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "stop"}));
});

if (window["WebSocket"]) {
    console.log("supports websockets");
    // Connect to websocket
    conn = new WebSocket("ws://" + document.location.host + "/ws");
} else {
    popupInfo.error("Not supporting websockets");
}

conn.addEventListener("open", () => {
    popupInfo.info("connected", 1);
});

conn.addEventListener("close", () => {
    console.log("disconnected");
});

conn.addEventListener("message", (evt) => {
    resp = JSON.parse(evt.data);
    popupInfo.info(resp);
});

