import express from "express";
import addRequestId from "express-request-id";
import morgan from "morgan";
import { MongoClient } from 'mongodb';

morgan.token("id", (req) => req.id.split("-")[0]);

const app = express();
let db = null;

MongoClient.connect('mongodb://localhost:27017', (err, client) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }
    db = client.db('otus');
});

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

app.get("/courses", (req, res) => {
    const cursor = db.collection('courses').find({}, {name: 1});
    cursor.toArray((err, courses) => {
        if (err) {
            console.log(err);
            res.sendStatus(500);
        } else {
            res.render("courses", { courses });
        }
    });
});

app.get("/courses/:name", (req, res) => {
    db.collection('courses').findOne({name: req.params.name}, (err, course) => {
        if (err) {
            console.log(err);
            res.sendStatus(500);
        } else if (!course) {
            res.sendStatus(404);
        } else {
            res.render("course", { course });
        }
    });
});

app.get("/test", (req, res) => {
    res.send("Hello world! You've sent me - " + req.query.name);
});

app.listen(3000, () => console.log("Hello from 3000"));
