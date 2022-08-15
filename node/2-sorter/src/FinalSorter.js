import fs from "fs";

import { Reader } from './Reader.js';

export class FinalSorter {
    constructor(fileNames, outputFileName) {
        this._fileNames = fileNames;
        this._outputFileName = outputFileName;
        this._readers = [];
        this._outputStream = fs.createWriteStream(outputFileName);
        this._outputNumbersCounter = 0;
    }
    
    sortFinally() {
        return new Promise((resolve, reject) => {
            this._resolve = resolve;
            this._reject = reject;
            
            this._readers = this._fileNames.map((fileName) => new Reader(fileName));
            Promise.all(this._readers.map((reader) => reader.start())).then(() => {
                this._iterateWithSorting();
            });
        });
    }
    
    _iterateWithSorting() {
        if (this._areReadersFinished()) {
            this._outputStream.end('', () => {
                this._resolve(this._outputNumbersCounter);
            });
            return;
        }
        
        const readerWithMinNumber = this._findReaderWithMinNumber();
        
        this._writeNumber(readerWithMinNumber.getNumber());
        
        readerWithMinNumber.loadNextNumber().then(() => {
            this._iterateWithSorting();
        });
    }
    
    _areReadersFinished() {
        return this._readers.every((reader) => reader.isFinished());
    }
    
    _findReaderWithMinNumber() {
        let minNumber = null;
        let minNumberReader = null;
        
        this._readers.forEach((reader) => {
            if (!reader.isFinished() && (minNumber === null || reader.getNumber() < minNumber)) {
                minNumber = reader.getNumber();
                minNumberReader = reader;
            }
        });
        
        return minNumberReader;
    }
    
    _writeNumber(number) {
        this._outputNumbersCounter++;
        this._outputStream.write(`${number} `);
    }
}