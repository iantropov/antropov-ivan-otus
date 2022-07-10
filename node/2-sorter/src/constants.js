import { absPath } from "./utils.js";

export const INPUT_FILE_NAME = absPath("/data/input.txt");
export const OUTPUT_FILE_NAME = absPath("/data/output.txt");
export const NUMBERS_COUNT = 15_000_000;
export const NUMBERS_COUNT_IN_PRE_SORT_FILE = Math.floor(NUMBERS_COUNT / 10);
