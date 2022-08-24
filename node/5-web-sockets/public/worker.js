const socket = new WebSocket("ws://localhost:8080");

socket.addEventListener("open", (event) => {
    console.log("[WORKER]: Made connection with the server");
    socket.send("Hello Server!");
});

socket.addEventListener("message", (event) => {
    console.log(`[WORKER]: Received a message from server: ${event.data}`);
    postMessage(event.data);
});

socket.addEventListener("close", () => {
    console.log("[WORKER]: Closed connection with server");
});

onmessage = (event) => {
    console.log(`[WORKER]: Received new message from page: ${event.data}`);
    socket.send(event.data);
};
