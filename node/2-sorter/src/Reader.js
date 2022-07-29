import { EventEmitter } from 'events'; 

import { NumberFileReadStream } from "./NumberFileReadStream.js";
import { logDebug } from './utils.js';

export class Reader extends EventEmitter {
    constructor(fileName) {
        super();
        
        this._fileName = fileName;
        this._isFinished = false;
        this._number = null;
        this._readableEventFired = false;
        this._readCount = 0;

        this._fileStream = new NumberFileReadStream(fileName);
        this._fileStream.on('end', this._onFileStreamEnd.bind(this));
        this._fileStream.on('readable', this._onFileStreamReadable.bind(this));
    }
    
    _onFileStreamEnd() {
        logDebug(`Reader:${this._fileName}: Finished. Read ${this._readCount} numbers.`);
    }

    _onFileStreamReadable() {
        logDebug(`Reader:${this._fileName}: Logged readable.`);
        
        this._readableEventFired = true;
        this.emit('streamIsReadable');
    }    
    
    start() {
        return new Promise((resolve, reject) => {
            if (this._readableEventFired) {
                this._number = this._fileStream.read();
                this._readCount++;
                logDebug(`Reader:${this._fileName}: Read number (start)(buffered): ${this._number}`);
                this._readableEventFired = false;
                resolve();
            } else {
                this.once('streamIsReadable', () => {
                    this._number = this._fileStream.read();
                    this._readCount++;
                    logDebug(`Reader:${this._fileName}: Read number (start)(on_event): ${this._number}`);
                    this._readableEventFired = false;
                    resolve();
                });
            }
        });
    }

    isFinished() {
        return this._isFinished;
    }

    getNumber() {
        return this._number;
    }

    loadNextNumber() {
        return new Promise((resolve, reject) => {
            const result = this._fileStream.read();
            if (result === null && this._readableEventFired) {
                logDebug(`Reader:${this._fileName}: Found already finished stream. Read ${this._readCount} numbers.`);
                this._isFinished = true;
                resolve();
            } else if (result === null) {
                logDebug(`Reader:${this._fileName}: Will wait (readable)`);
                this.once("streamIsReadable", () => {
                    logDebug(`Reader:${this._fileName}: Emitted (readable)`);
                    const result = this._fileStream.read();
                    this._readableEventFired = false;
                    if (result === null) {
                        this._isFinished = true;
                    } else {
                        logDebug(`Reader:${this._fileName}: Read number (on_event): ${result}`);
                        this._number = result;
                        this._readCount++;
                    }
                    resolve();
                });
            } else {
                this._readableEventFired = false;
                logDebug(`Reader:${this._fileName}: Read number(buffered): ${result}`);
                this._number = result;
                this._readCount++;
                resolve();
            }
        });
    }
}
