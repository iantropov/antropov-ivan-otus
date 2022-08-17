import mongoose from "mongoose";

const lessonSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,
    },
    description: {
        type: String,
        required: true,
    },
});

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
