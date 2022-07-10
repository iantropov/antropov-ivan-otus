import { NumberFileGenerator } from "./NumberFileGenerator.js";
import { logInfo, fileSize, memoryUsageInMb } from "./utils.js";
import { INPUT_FILE_NAME, NUMBERS_COUNT } from "./constants.js";

console.time("generationTime");
logInfo("Starting generating numbers...");
const generator = new NumberFileGenerator(INPUT_FILE_NAME, NUMBERS_COUNT);
generator.generate().then((numbersCount) => {
    logInfo(
        `Generated ${numbersCount} numbers. File size: ${fileSize(
            INPUT_FILE_NAME
        )} MB`
    );
    console.timeEnd("generationTime");
    console.log(`The script uses approximately ${memoryUsageInMb()} MB`);
});
