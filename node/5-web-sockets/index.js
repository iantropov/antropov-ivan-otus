import express from "express";
import { WebSocketServer } from "ws";

const wss = new WebSocketServer({ port: 8080 });

let lastClientWs = null;
let message = "initial message";

wss.on("connection", (ws) => {
    lastClientWs = ws;

    console.log("Received WS connection!");

    ws.on("message", (data) => {
        console.log("received: %s", data);
        message = data.toString();

        if (message === 'finish') {
            lastClientWs = null;
            ws.close();
        }
    });

    ws.send("something");
});

const app = express();

app.use(express.static("public"));

setInterval(() => {
    if (lastClientWs) {
        console.log(`Sending ${message} to client`);
        lastClientWs.send(message);
    }
}, 5000);

app.listen(3000, () => console.log("Hello from 3000"));
