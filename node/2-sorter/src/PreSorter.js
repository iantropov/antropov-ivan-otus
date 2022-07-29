import fs from "fs";
import util from "util";

import { NumberFileReadStream } from "./NumberFileReadStream.js";
import { absPath, logInfo, logError, logDebug } from "./utils.js";

const writeFile = util.promisify(fs.writeFile);

export class PreSorter {
    constructor(fileName, preSortBufferSize) {
        this._inputStream = new NumberFileReadStream(fileName);
        this._preSortBufferSize = preSortBufferSize;
        this._outputBuffer = [];
        this._outputFiles = [];
    }

    preSort() {
        return new Promise((resolve, reject) => {
            this._inputStream.on("data", this._onInputStreamData.bind(this));
            this._inputStream.on("end", this._onInputStreamEnd.bind(this));

            this._presortResolve = resolve;
            this._presortReject = reject;
        });
    }
    
    removePreSortFiles() {
        this._outputFiles.forEach(({ fileName }) => {
            fs.rm(fileName, () => {});
        });
    }

    _buildFileEntry(fileName, writeFilePromise) {
        return {
            fileName,
            writeFilePromise
        };
    }

    _onInputStreamData(number) {
        logDebug('PreSorter: onData');
        this._outputBuffer.push(number);
        if (this._outputBuffer.length === this._preSortBufferSize) {
            this._dumpNextOutputFile();
        }
    }

    _onInputStreamEnd() {
        if (this._outputBuffer.length > 0) {
            this._dumpNextOutputFile();
        }

        Promise.all(
            this._outputFiles.map(({ writeFilePromise }) => writeFilePromise)
        ).then(
            () => {
                this._presortResolve(
                    this._outputFiles.map(({ fileName }) => fileName)
                );
            },
            (error) => {
                this._presortReject(error);
            }
        );
    }

    _dumpNextOutputFile() {
        const nextOutputFileName = this._outputFileName(this._outputFiles.length);

        const outFilePromise = this._saveOutputFile(
            nextOutputFileName,
            this._outputBuffer.sort((a, b) => a - b).join(" ")
        );

        this._outputFiles.push(
            this._buildFileEntry(nextOutputFileName, outFilePromise)
        );

        this._outputBuffer = [];
    }
    
    _outputFileName(index) {
        return absPath(`/data/tmp.${index}.txt`);
    }
    
    _saveOutputFile(name, content) {
        return writeFile(name, content).then(
            () => {
                logInfo(`Successfull saved file: ${name}`);
            },
            (error) => {
                logError(error);
            }
        );
    }
}
