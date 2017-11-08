#!/bin/env/node
const express = require('express');
const http = require('http');
const path = require('path');
const pugToHtml = require('./pugToHtml');

const config = require('./config');
const app = express();
const templateDir = path.join(config.appRoot, 'templates');
const publicDir = path.join(config.appRoot, 'public');
const scriptsDir = path.join(config.appRoot, 'scripts');

// app config
app.set('port', config.port);
app.use(express.static(publicDir));
app.use('/scripts', express.static(scriptsDir));

// initialize~
pugToHtml(templateDir, publicDir);
http.createServer(app).listen(app.get('port'), function () {
    console.log('Big Zam listening on port %d...', app.get('port'));
});

process.on('SIGINT', function() {
    process.exit();
});

