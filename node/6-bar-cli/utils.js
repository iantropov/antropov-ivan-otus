const BAR_OWNERS_CREDENTIALS = 'iam:boss'

module.exports.authorizeUser = function authorizeUser(options) {
    console.dir(options);
    if (!options.credentials) {
        console.log("You must supply your credentials in order to change drinks");
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
