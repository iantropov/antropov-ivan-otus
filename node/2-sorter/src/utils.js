import path from "node:path";
import fs from "node:fs";

export function absPath(relativePath) {
    return path.join(process.cwd(), relativePath);
}

export function logDebug(message) {
    if (process.env.DEBUG) {
        console.debug(message);
    }
}

export function fileSize(path) {
    const size = fs.statSync(path).size;
    const sizeInMb = size / 1024 / 1024
    return Math.floor(sizeInMb * 10_000) / 10_000;
}

export function logInfo(message) {
    console.info(message);
}

export function logError(message) {
    console.error(message);
}

export function memoryUsageInMb() {
    const used = process.memoryUsage().heapUsed / 1024 / 1024;
    return Math.round(used * 100) / 100
}