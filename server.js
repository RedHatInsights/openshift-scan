#!/bin/env/node
const express = require('express');
const http = require('http');
const path = require('path');
const bodyParser = require('body-parser');
const pugToHtml = require('./pugToHtml');

const app = express();
const appRoot = path.join(__dirname, 'app');
const templateDir = path.join(appRoot, 'templates');
const publicDir = path.join(appRoot, 'public');
const scriptsDir = path.join(appRoot, 'scripts');

// app config
app.set('port', 8080);
app.use(express.static(publicDir));
app.use('/scripts', express.static(scriptsDir));

// routes
app.post('/report', bodyParser.json(), (req, res) => {
    //
});

app.get('/report/:id', (req, res) => {
    //
});


// initialize~
pugToHtml(templateDir, publicDir);
http.createServer(app).listen(app.get('port'), function () {
    console.log('Big Zam listening on port %d...', app.get('port'));
});

process.on('SIGINT', function() {
    process.exit();
});

