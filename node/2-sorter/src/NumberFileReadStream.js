import { Readable } from "node:stream";
import fs from "node:fs";

import { logDebug, logError } from "./utils.js";

const BUFFER_SIZE = 1024;
const MAX_INTEGER_LENGTH_IN_BYTES = 16;
const MIN_CODE = 48;
const MAX_CODE = 57;

export class NumberFileReadStream extends Readable {
    constructor(fileName) {
        super({
            objectMode: true,
            highWaterMark: 1,
        });
        this._fileName = fileName;
        this._fileDescriptor = null;
        this._rawBuffer = Buffer.alloc(BUFFER_SIZE);
        this._numberBuffer = Buffer.alloc(MAX_INTEGER_LENGTH_IN_BYTES);
        this._numberBufferLength = 0;
        this._readyNumbers = [];
        this._isFinishedRead = false;
    }

    _construct(callback) {
        fs.open(this._fileName, (error, fd) => {
            if (error) {
                logError(error);
                callback(error);
            } else {
                this._fileDescriptor = fd;
                callback();
            }
        });
    }

    _read() {
        if (this._pushReadyNumbers()) {
            return;
        }

        fs.read(
            this._fileDescriptor,
            this._rawBuffer,
            0,
            BUFFER_SIZE,
            null,
            (error, bytesRead) => {
                if (error) {
                    logError(error);
                    this.destroy(error);
                } else {
                    if (bytesRead === 0) {
                        logDebug(
                            `NumberFileReadStream:${this._fileName}: finished read.`
                        );
                        this._isFinishedRead = true;
                        if (this._numberBufferLength > 0) {
                            this._extractNumber();
                        }
                        this._pushReadyNumbers();
                    } else {
                        this._parseRawBuffer(bytesRead);

                        if (
                            bytesRead < BUFFER_SIZE &&
                            this._numberBufferLength > 0
                        ) {
                            this._extractNumber();
                        }

                        if (this._readyNumbers.length > 0) {
                            this._pushReadyNumbers();
                        } else {
                            this._read();
                        }
                    }
                }
            }
        );
    }

    _parseRawBuffer(bytesRead) {
        for (let i = 0; i < bytesRead; i++) {
            if (
                this._rawBuffer[i] >= MIN_CODE &&
                this._rawBuffer[i] <= MAX_CODE
            ) {
                this._numberBuffer[this._numberBufferLength++] =
                    this._rawBuffer[i];
            } else if (this._numberBufferLength > 0) {
                this._extractNumber();
            }
        }
    }

    _pushReadyNumbers() {
        if (this._readyNumbers.length > 0) {
            const nextNumber = this._readyNumbers.shift();
            logDebug(
                `NumberFileReadStream:${this._fileName}: pushed number ${nextNumber}.`
            );
            this.push(nextNumber);
            return true;
        } else if (this._isFinishedRead) {
            logDebug(`NumberFileReadStream:${this._fileName}: pushed null.`);
            this.push(null);
            return true;
        }
    }

    _extractNumber() {
        this._readyNumbers.push(
            Number(
                this._numberBuffer.toString(
                    "utf-8",
                    0,
                    this._numberBufferLength
                )
            )
        );
        this._numberBufferLength = 0;
    }

    _destroy(err, callback) {
        if (this._fileDescriptor) {
            fs.close(this._fileDescriptor, (er) => callback(er || err));
        } else {
            callback(err);
        }
    }
}
