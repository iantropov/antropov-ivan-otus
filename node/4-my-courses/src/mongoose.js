import mongoose from "mongoose";

export default {
    start() {
        console.log("Connecting to MongoDB...")
        return mongoose.connect("mongodb://localhost:27017").then(() => {
            console.log("Successfully connected to MongoDB");
        }, (err) => {
            console.error("Failed to connect to MongoDB", err);
        });
    }
}
