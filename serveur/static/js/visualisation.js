const btnStart = document.getElementById("start-simu");
const btnPause = document.getElementById("pause-simu");
const btnContinue = document.getElementById("continue-simu");
const btnStop = document.getElementById("stop-simu");

const valTest = document.getElementById("value-test");
let conn = undefined;

const currentURL = window.location.href;
const url = new URL(currentURL);
const pathname = url.pathname;
const pathParts = pathname.split('/');
const id = pathParts[pathParts.length - 1];
console.log("ID extrait de l'URL :", id);


//ENVOYER ICI EN PRENANT window.location.href ????

btnStart.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "start"}));
    //btnStart.disabled = true
});

btnPause.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "pause"}));
});

btnContinue.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "continue"}));
});

btnStop.addEventListener("click", () => {
    conn.send(JSON.stringify({id_simulation : id, type: "action", data: "stop"}));
});

if (window["WebSocket"]) {
    console.log("supports websockets");
    // Connect to websocket
    conn = new WebSocket("ws://" + document.location.host + "/ws");
} else {
    alert("Not supporting websockets");
}

conn.addEventListener("open", () => {
    console.log("connected !!!");
});

conn.addEventListener("close", () => {
    console.log("disconnected !!!");
});

conn.addEventListener("message", (evt) => {
    resp = JSON.parse(evt.data)
    valTest.innerText = resp;
})