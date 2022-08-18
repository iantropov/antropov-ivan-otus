import express from "express";
import addRequestId from "express-request-id";
import morgan from "morgan";
import session from "express-session";

import mongoose from "./mongoose.js";
import { registerCoursesRoutes } from "./routes/courses.js";
import { registerLoginRoutes } from "./routes/login.js";

morgan.token("id", (req) => req.id.split("-")[0]);

const app = express();

app.set("view engine", "pug");
app.set("views", "./src/views");

app.use(addRequestId({ setHeader: false }));

app.use(
    session({
        secret: "90eHRm0MgAADMq_kbNrLAg",
        cookie: { maxAge: 60000 },
        saveUninitialized: false,
        resave: false,
    })
);

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

app.use((req, res, next) => {
    if (req.path !== "/login" && !req.session.userName) {
        res.redirect("/login");
    } else {
        next();
    }
});

mongoose.start().then(() => {
    registerLoginRoutes(app);
    registerCoursesRoutes(app);
});

app.listen(3000, () => console.log("Hello from 3000"));
