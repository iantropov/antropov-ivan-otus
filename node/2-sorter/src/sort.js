import { PreSorter } from "./PreSorter.js";
import { FinalSorter } from "./FinalSorter.js";
import { logInfo, memoryUsageInMb } from "./utils.js";
import { INPUT_FILE_NAME, NUMBERS_COUNT_IN_PRE_SORT_FILE, OUTPUT_FILE_NAME } from "./constants.js";

console.time("totalTime");
console.time("presortingTime");
logInfo("Starting presorting....");
const preSorter = new PreSorter(INPUT_FILE_NAME, NUMBERS_COUNT_IN_PRE_SORT_FILE);
preSorter.preSort().then((outputFileNames) => {
    logInfo(`Finished presorting with ${outputFileNames.length} files.`);
    console.timeEnd("presortingTime");
    
    console.time("finalSortingTime");
    logInfo("Starting final sorting...");

    const finalSorter = new FinalSorter(outputFileNames, OUTPUT_FILE_NAME);
    finalSorter.sortFinally().then((sortedNumbersCount) => {
        logInfo(`Sorted ${sortedNumbersCount} numbers.`);
        console.timeEnd("finalSortingTime");

        preSorter.removePreSortFiles();
        logInfo(`Removed ${outputFileNames.length} preSort files.`);
        console.timeEnd("totalTime");
        
        console.log(`The script uses approximately ${memoryUsageInMb()} MB`);
    });
});

