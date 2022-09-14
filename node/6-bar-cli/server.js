const program = require("commander");
const { writeFileSync } = require('node:fs');

const drinks = require("./data/drinks.json");
const { name, version, description } = require("./package.json");

const BAR_OWNERS_CREDENTIALS = 'iam:boss'

program.name(name).description(description).version(version);

program
    .command("get-drinks")
    .description("Get a list of available drinks")
    .action(() => {
        console.log(drinks);
    });

program
    .command("get-drink")
    .argument("<name>", "Name of the drink")
    .action((drinkName) => {
        const drink = drinks.find(({ name }) => name === drinkName);
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
    .option("-v, --volume <volume>", "Volume of the drink", '1')
    .action((options) => {
        if (!options.credentials) {
            console.log("You must supply your credentials in order to add a drink");
            return;
        }

        if (options.credentials !== BAR_OWNERS_CREDENTIALS) {
            console.log("We can't allow you to add a drink. You aren't the boss!");
            return;
        }

        if (!options.name) {
            console.log("You must supply name of the drink");
            return;
        }

        const volumeNumber = Number(options.volume);

        const newDrink = {
            name: options.name,
            volume: isNaN(volumeNumber) ? 1 : volumeNumber
        }
        console.log('Here you are! This is your new drink: ', newDrink);

        drinks.push(newDrink)

        writeFileSync('./data/drinks.json', JSON.stringify(drinks));
    });


program.parse(process.argv);
