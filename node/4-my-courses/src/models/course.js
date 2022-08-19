import mongoose from "mongoose";

import { lessonSchema } from "./lesson.js";

const courseSchema = new mongoose.Schema(
    {
        name: {
            type: String,
            required: true,
            index: true,
        },
        lessons: {
            type: [lessonSchema],
            required: true,
        },
    },
    { autoIndex: true }
);

export const Course = mongoose.model("Course", courseSchema);
