import { Course } from "../models/course.js";

export const registerCoursesRoutes = (express) => {
    express.get("/courses", async (req, res) => {
        const courses = await Course.find();
        res.render("courses", { courses });
    });

    express.get("/courses/:courseName", async (req, res) => {
        const course = await Course.findOne({ name: req.params.courseName });
        if (!course) {
            res.sendStatus(404);
            return;
        }
        res.render("course", { course });
    });

    express.get("/courses/:courseName/lessons/:lessonName", async (req, res) => {
        const course = await Course.findOne({ name: req.params.courseName });
        if (!course) {
            res.sendStatus(404);
            return;
        }

        const lesson = course.lessons.find((lesson) => lesson.name === req.params.lessonName);
        if (!lesson) {
            res.sendStatus(404);
            return;
        }

        res.render("lesson", { lesson });
    });
};
