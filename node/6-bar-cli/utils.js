var Table = require("cli-table");

const BAR_OWNERS_CREDENTIALS = "iam:boss";

module.exports.authorizeUser = function authorizeUser(options) {
    console.dir(options);
    if (!options.credentials) {
        console.log(
            "You must supply your credentials in order to change drinks"
        );
        return false;
    }

    if (options.credentials !== BAR_OWNERS_CREDENTIALS) {
        console.log("We can't allow you to add a drink. You aren't the boss!");
        return false;
    }

    return true;
};

module.exports.extractDrink = function extractDrink(options) {
    if (!options.name) {
        console.log("You must supply name of the drink");
        return null;
    }

    const volumeNumber = Number(options.volume);

    return {
        name: options.name,
        volume: isNaN(volumeNumber) ? 1 : volumeNumber,
    };
};

module.exports.printDrinks = function printDrinks(drinks) {
    var table = new Table({
        head: ["Index", "Name", "Volume"],
        colWidths: [10, 20, 10],
        rows: drinks.map((drink, index) => [index + 1, drink.name, drink.volume]),
    });

    console.log(table.toString());
};
