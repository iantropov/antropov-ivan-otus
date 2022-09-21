import { readFileSync, writeFileSync } from "node:fs";

const DRINKS_PATH = "./data/drinks.json";
const DEFAULT_DRINKS = [
    { name: "cola", volume: 0.5 },
    { name: "old-drink", volume: 6 },
];

let drinks = null;
try {
    drinks = JSON.parse(readFileSync(DRINKS_PATH, "utf-8"));
} catch (error) {
    if (error.code === "ENOENT") {
        drinks = DEFAULT_DRINKS;
    }
}

export function saveDrinks() {
    writeFileSync(DRINKS_PATH, JSON.stringify(drinks));
}

export function findAll() {
    return drinks;
}

export function findDrink(drinkName) {
    return drinks.find(({ name }) => name === drinkName);
}

export function findDrinkIndex(drinkName) {
    return drinks.findIndex(({ name }) => name === drinkName);
}

export function addDrink(drink) {
    drinks.push(drink);
}

export function deleteDrink(drinkName) {
    const indexToDelete = this.findDrinkIndex(drinkName);
    drinks.splice(indexToDelete, 1);
}
