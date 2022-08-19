import mongoose from "mongoose";

export const lessonSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,
    },
    description: {
        type: String,
        required: true,
    },
});

export const Lesson = mongoose.model("Lesson", lessonSchema);
