const { readFileSync, writeFileSync } = require("node:fs");

const DRINKS_PATH = "./data/drinks.json";
const DEFAULT_DRINKS = [{"name":"cola","volume":0.5},{"name":"old-drink","volume":6}];

let drinks = null;
try {
    drinks = JSON.parse(readFileSync(DRINKS_PATH, 'utf-8'));
} catch(error) {
    if (error.code === 'ENOENT') {
        drinks = DEFAULT_DRINKS;
    }
}

module.exports = {
    saveDrinks() {
        writeFileSync(DRINKS_PATH, JSON.stringify(drinks));
    },

    findAll() {
        return drinks;
    },

    findDrink(drinkName) {
        return drinks.find(({ name }) => name === drinkName);
    },

    findDrinkIndex(drinkName) {
        return drinks.findIndex(({ name }) => name === drinkName);
    },

    addDrink(drink) {
        drinks.push(drink);
    },

    deleteDrink(drinkName) {
        const indexToDelete = this.findDrinkIndex(drinkName);
        drinks.splice(indexToDelete, 1);
    },
};
