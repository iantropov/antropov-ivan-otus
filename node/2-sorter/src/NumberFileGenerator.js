import fs from "fs";

const CHUNK_SIZE = 250_000;

export class NumberFileGenerator {
    constructor(fileName, numbersCount) {
        this._outputStream = fs.createWriteStream(fileName);
        this._numbersCount = numbersCount;
        this._allNumbersCount = numbersCount;
    }

    generate() {
        return new Promise((resolve, reject) => {
            this._resolveGeneration = resolve;
            this._rejectGeneration = reject;
            
            this._generateChunk();
        });
    }

    _generateChunk() {
        const { chunk, numbersCountInChunk } = this._buildChunk();
        this._outputStream.write(chunk, "utf-8", (err) => {
            debugger
            if (err) {
                console.error(err);
                this._rejectGeneration(err);
            }

            this._numbersCount -= numbersCountInChunk;
            if (this._numbersCount > 0) {
                this._generateChunk();
            } else {
                this._resolveGeneration();
            }
        });
    }

    _buildChunk() {
        const numbersCountInChunk = Math.min(this._numbersCount, CHUNK_SIZE);
        const array = new Uint32Array(numbersCountInChunk);
        for (let i = 0; i < numbersCountInChunk; i++) {
            array[i] = Math.floor(Math.random() * this._numbersCount);
        }
        return {
            chunk: array.join(" ") + " ",
            numbersCountInChunk
        };
    }
}
