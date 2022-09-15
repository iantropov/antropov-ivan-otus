const { writeFileSync } = require("node:fs");

const DRINKS_PATH = "./data/drinks.json";

const drinks = require(DRINKS_PATH);

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
