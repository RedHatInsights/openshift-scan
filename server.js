#!/bin/env/node
const express = require('express');
const http = require('http');
const path = require('path');
const pugToHtml = require('./pugToHtml');

const app = express();
const appRoot = path.join(__dirname, 'app');
const templateDir = path.join(appRoot, 'templates');
const publicDir = path.join(appRoot, 'public');

app.set('port', 8081);
app.use(express.static(publicDir));

pugToHtml(templateDir, publicDir);

http.createServer(app).listen(app.get('port'), function () {
    console.log('Big Zam listening on port %d...', app.get('port'));
});

