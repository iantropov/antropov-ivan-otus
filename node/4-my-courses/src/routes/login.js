import bodyParser from "body-parser";
import csrf from "csurf";

const parseForm = bodyParser.urlencoded({ extended: false });
const csrfProtection = csrf();

export const registerLoginRoutes = (express) => {
    express.get("/login", csrfProtection, async (req, res) => {
        res.render("login", { csrfToken: req.csrfToken() });
    });

    express.post("/login", parseForm, csrfProtection, async (req, res) => {
        if (req.body.password !== "password") {
            res.sendStatus(401);
        } else {
            req.session.userName = req.body.name;
            res.redirect("/courses");
        }
    });

    express.use(function (err, req, res, next) {
        if (err.code !== "EBADCSRFTOKEN") return next(err);
        res.sendStatus(401);
    });
};
