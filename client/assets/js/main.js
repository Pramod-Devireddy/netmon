const websocket = new ReconnectingWebSocket("ws://127.0.0.1:9001/ws");

websocket.onmessage = function (event) {
    // parse the event data sent from our websocket server
    var subscribers = JSON.parse(event.data);
    // populate our `sub` element with the total subscriber counter for our
    // channel
    document.getElementById("subs").innerText = subscribers;
};