const env = process.env;
const path = require('path');

config = {
    port: 8080,
    root: path.normalize(__dirname + '/..'),
    appRoot: path.normalize(__dirname + '/../app'),
};

module.exports = config;