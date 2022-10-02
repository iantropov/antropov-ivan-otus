import Table from "cli-table";
import { Drink } from "drinks-model";
import { createRequire } from "module";

const BAR_OWNERS_CREDENTIALS = "iam:boss";

export interface BarCliOptions {
    credentials?: string;
    name?: string;
    volume?: string;
}

export function loadPackageJson() {
    const require = createRequire(import.meta.url);
    return require("../package.json");
}

export function authorizeUser(options: BarCliOptions) {
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
}

export function extractDrink(options: BarCliOptions): Drink | null {
    if (!options.name) {
        console.log("You must supply name of the drink");
        return null;
    }

    const volumeNumber = Number(options.volume);

    return {
        name: options.name,
        volume: isNaN(volumeNumber) ? 1 : volumeNumber,
    };
}

export function printDrinks(drinks: Drink[]) {
    const table = new Table({
        head: ["#", "Name", "Volume"],
        colWidths: [10, 20, 10],
    });

    drinks.forEach((drink, index) => {
        table.push([index + 1, drink.name, drink.volume]);
    });

    console.log(table.toString());
}
