import { readFileSync, writeFileSync } from "node:fs";

export interface Drink {
    name: string;
    volume: number;
}

const DRINKS_PATH = "./data/drinks.json";
const DEFAULT_DRINKS: Drink[] = [
    { name: "cola", volume: 0.5 },
    { name: "old-drink", volume: 6 },
];

let drinks: Drink[] = [];
try {
    drinks = JSON.parse(readFileSync(DRINKS_PATH, "utf-8"));
} catch (error: unknown) {
    if ((error as NodeJS.ErrnoException)?.code === "ENOENT") {
        drinks = DEFAULT_DRINKS;
    }
}

export function saveDrinks() {
    writeFileSync(DRINKS_PATH, JSON.stringify(drinks));
}

export function findAll() {
    return drinks;
}

export function findDrink(drinkName: Drink['name']): Drink | undefined {
    return drinks.find(({ name }) => name === drinkName);
}

export function findDrinkIndex(drinkName: Drink['name']) {
    return drinks.findIndex(({ name }) => name === drinkName);
}

export function addDrink(drink: Drink) {
    drinks.push(drink);
}

export function deleteDrink(drinkName: Drink['name']) {
    const indexToDelete = findDrinkIndex(drinkName);
    drinks.splice(indexToDelete, 1);
}
