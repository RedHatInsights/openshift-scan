const gulp = require('gulp');
const nodemon = require('gulp-nodemon');
const browserSync = require('browser-sync').create();
const config = require('./config');
 
gulp.task('default', ['bsync'], () => {});

gulp.task('bsync', ['nodemon'], () => {
    browserSync.init(null, {
        proxy: 'http://localhost:' + config.port,
        files: ['app/**/*/*.*'],
        port: 7000
    });
});

gulp.task('nodemon', ['initfiles'], (done) => {
    let started = false;
    return nodemon({
        script: 'server.js',
        ext: 'js jade json'
    }).on('start', () => {
        if (!started) {
            setTimeout(function() {
                started = true;
                done();
            }, 3000);

        }
    });
});
