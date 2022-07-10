import fs from "fs";

export class NumberFileGenerator {
    constructor(fileName, numbersCount) {
        this._outputStream = fs.createWriteStream(fileName);
        this._numbersCount = numbersCount;
    }

    generate() {
        return new Promise((resolve) => {
            for (let i = 0; i < this._numbersCount; i++) {
                this._outputStream.write(
                    `${Math.floor(Math.random() * this._numbersCount)} `
                );
            }
            this._outputStream.end("", () => {
                resolve(this._numbersCount);
            });
        });
    }
}
