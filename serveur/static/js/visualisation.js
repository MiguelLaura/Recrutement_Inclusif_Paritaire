const btnStart = document.getElementById("launch-simu");
const valTest = document.getElementById("value-test");
let conn = undefined;

btnStart.addEventListener("click", () => {
    conn.send(JSON.stringify({type: "salutation", data: "salut le serveur :)"}));
    btnStart.disabled = true
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
    console.log(evt, JSON.parse(evt.data));
})