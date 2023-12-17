// -------------------------
// WebSocket

let conn = undefined;

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
});