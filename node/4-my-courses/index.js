import express from "express";
import addRequestId from "express-request-id";
import morgan from "morgan";

morgan.token("id", (req) => req.id.split("-")[0]);

const app = express();

app.set("view engine", "pug");

app.use(addRequestId({ setHeader: false }));

app.use(
    morgan("[:date[iso] #:id] Started :method :url for :remote-addr", {
        immediate: true,
    })
);
app.use(
    morgan(
        "[:date[iso] #:id] Completed :status :res[content-length] in :response-time ms"
    )
);

// app.use((req, res, next) => {
//     console.log(`[${new Date().toISOString()}] - Request received`);
//     next();
// });

app.get("/", (req, res) => {
    res.render("index", { title: "Hello title", message: "Hello message" });
});

app.get("/test", (req, res) => {
    res.send("Hello world! You've sent me - " + req.query.name);
});

app.listen(3000, () => console.log("Hello from 3000"));
