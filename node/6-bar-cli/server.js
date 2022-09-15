const program = require("commander");

const { name, version, description } = require("./package.json");
const drinksModel = require("./drinks-model.js");
const { authorizeUser, extractDrink } = require("./utils.js");

program.name(name).description(description).version(version);

program
    .command("get-drinks")
    .description("Get a list of available drinks")
    .action(() => {
        console.log(drinksModel.findAll());
    });

program
    .command("get-drink")
    .argument("<name>", "Name of the drink")
    .action((drinkName) => {
        const drink = drinksModel.findDrink(drinkName);
        if (!drink) {
            console.log("Oops, we don't have this drink.");
        } else {
            console.log(drink);
        }
    });

program
    .command("add-drink")
    .option("-c, --credentials <name:password>", "Credentials of the bar owner")
    .option("-n, --name <name>", "Name of the drink")
    .option("-v, --volume <volume>", "Volume of the drink", "1")
    .action((options) => {
        if (!authorizeUser(options)) {
            return;
        }

        const drinkFromOptions = extractDrink(options);
        if (!drinkFromOptions) {
            return;
        }

        console.log("Here you are! This is your new drink: ", drinkFromOptions);

        drinksModel.addDrink(drinkFromOptions);
        drinksModel.saveDrinks();
    });

program
    .command("update-drink")
    .argument("<name>", "Name of the drink to update")
    .option("-c, --credentials <name:password>", "Credentials of the bar owner")
    .option("-n, --name <name>", "Name of the drink")
    .option("-v, --volume <volume>", "Volume of the drink", "1")
    .action((drinkName, options) => {
        if (!authorizeUser(options)) {
            return;
        }

        const drinkToChange = drinksModel.findDrink(drinkName);
        if (!drinkToChange) {
            console.log("Oops, we don't have this drink.");
            return;
        }

        const drinkFromOptions = extractDrink(options);
        if (!drinkFromOptions) {
            return;
        }

        drinkToChange.name = drinkFromOptions.name;
        drinkToChange.volume = drinkFromOptions.volume;

        console.log(
            "Here you are! This is your updated drink: ",
            drinkToChange
        );

        drinksModel.saveDrinks();
    });

program
    .command("delete-drink")
    .argument("<name>", "Name of the drink to update")
    .option("-c, --credentials <name:password>", "Credentials of the bar owner")
    .action((drinkName, options) => {
        if (!authorizeUser(options)) {
            return;
        }

        const drinkToDelete = drinksModel.findDrink(drinkName);
        if (!drinkToDelete) {
            console.log("Oops, we don't have this drink.");
            return;
        }

        drinksModel.deleteDrink(drinkName);

        console.log(
            "Here you are! These are your remain drinks: ",
            drinksModel.findAll()
        );

        drinksModel.saveDrinks();
    });

program.parse(process.argv);
