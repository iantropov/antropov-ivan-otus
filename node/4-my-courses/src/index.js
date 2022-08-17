import express from "express";
import addRequestId from "express-request-id";
import morgan from "morgan";

import mongoose from "./mongoose.js";
import { registerCoursesRoutes } from "./routes/courses.js";

morgan.token("id", (req) => req.id.split("-")[0]);

const app = express();

app.set("view engine", "pug");
app.set("views", "./src/views");

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

mongoose.start().then(() => {
    registerCoursesRoutes(app);
});

app.listen(3000, () => console.log("Hello from 3000"));
