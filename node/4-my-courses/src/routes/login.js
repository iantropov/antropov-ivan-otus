import bodyParser from 'body-parser';

export const registerLoginRoutes = (express) => {
    express.use(bodyParser.urlencoded({ extended: false }));

    express.get("/login", async (req, res) => {
        res.render("login");
    });

    express.post("/login", async (req, res) => {
        if (req.body.password !== 'password') {
            res.sendStatus(403);
        } else {
            req.session.userName = req.body.name;
            res.redirect('/courses');
        }
    });
};
