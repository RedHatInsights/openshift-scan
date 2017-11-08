const fs = require('fs');
const path = require('path');
const renderFile = require('pug').renderFile;

module.exports = function compilePug(pugDir, htmlDir) {
    const pugFiles = fs.readdirSync(pugDir);
    pugFiles.forEach(p => {
        console.log('Compiling %s...', p);
        const htmlFile = path.join(htmlDir, p.replace(path.extname(p), '.html'));
        const htmlContent = renderFile(path.join(pugDir, p));
        fs.writeFile(htmlFile, htmlContent, err => {
            if (err) {
                console.log('Error compiling %s: %s', p, err);
                process.exit(1);
            }
        });
    });
};
