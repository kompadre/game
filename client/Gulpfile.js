/**
 * Created by Alex on 27/11/2014.
 */


const gulp = require('gulp');
const clean = require('gulp-clean');
const concat = require('gulp-concat');
const notifier = require("node-notifier");
const connect = require('gulp-connect');
const sass = require('gulp-sass');
const rename = require('gulp-rename');
const event_stream = require('event-stream');

var gutil = require('gulp-util');
var through = require('through2');
var XLSX = require('xlsx');
var File = require('vinyl');

const fs = require('fs');

const html = 'html/index.html';

gulp.task('clean', function () {
    return gulp.src('public', { read: false })
        .pipe(clean());
});

gulp.task('html', function () {
    return gulp.src(html)
        .pipe(gulp.dest('public/'))
        .pipe(connect.reload());
});

const styleGlob = 'src/styles/**/*.css';
const sassGlob = 'src/styles/**/*.scss';

gulp.task('css', function () {
    return event_stream.merge(
        gulp.src(styleGlob),
        gulp.src(sassGlob)
            .pipe(sass().on('error', sass.logError))
    )
        .pipe(concat('main.css'))
        .pipe(gulp.dest('public/css/'))
        .pipe(connect.reload());

});

const localizationGlob = 'assets/database/text/data.xlsx';

gulp.task('watch::localization', () => {
    return gulp.watch(localizationGlob, watchSettings, gulp.parallel('localization'));
});

gulp.task('localization', () => {
    const destination = 'database/text/';

    const converter = through.obj(function (file, enc, cb) {
        var task = this;
        if (file.isNull()) {
            this.push(file);
            return cb();
        }

        if (file.isStream()) {
            this.emit('error', new gutil.PluginError('Locale Lazy Kitty', 'Streaming not supported'));
            return cb();
        }


        /* Call XLSX */
        var workbook = XLSX.read(file.contents, { type: "buffer" });

        var worksheet = workbook.Sheets[workbook.SheetNames[0]];

        const j = XLSX.utils.sheet_to_json(worksheet, { raw: false, header: 1 });

        const headerRow = j[0];

        const languages = headerRow.slice(1);

        const rowCount = j.length;

        for (let i = 0; i < languages.length; i++) {
            const language = languages[i];

            //generate a key-value file
            var json = {};

            for (let k = 1; k < rowCount; k++) {
                const row = j[k];

                const key = row[0];
                const value = row[i + 1];

                if (value === undefined) {
                    //no value
                    continue;
                }


                //replace special characters
                const formattedValue = value.replace(/\&\#([a-fA-F0-9]+)\;/gi, function (match, code) {
                    return String.fromCharCode(code);
                });

                if (json.hasOwnProperty(key)) {

                    if (json[key] !== formattedValue) {
                        //re-definition
                        console.error(`duplicate key definition:'${key}', old value='${json[key]}', new value='${value}', keeping old value`);
                    }

                    continue;
                }

                json[key] = formattedValue;
            }


            const targetFile = new File({
                cwd: '.',
                path: destination + language + '.json', // put each translation file in a folder
                contents: Buffer.from(JSON.stringify(json, null, 4)),
            });

            task.push(targetFile);

            console.log("Written file: " + file.path + " => " + targetFile.path);

        }


        cb();
    });

    return gulp.src(localizationGlob)
        .pipe(converter)
        .pipe(gulp.dest(dataSourcePath))
});


let dataTypesForCopy = ["json", "js"]
    .concat(["glb", "gltf"]) //3d models
    .concat(["png", "jpeg", "jpg", "svg", "dds", "bmp"])
    .concat(["ogg", "mp3", "wav"])
    .concat(["ttf"]) //fonts
    .concat(["bin"]) //binary data
;

const dataSourcePath = "app/data";
const dataTargetPath = './public/data/';

const dataSlug = dataSourcePath + "/**/*.{" + dataTypesForCopy.concat(dataTypesForCopy.map(t => t.toLocaleUpperCase())).join(',') + "}";

gulp.task('copy-data', function () {
    return gulp.src(dataSlug)
        .pipe(rename((path) => {
            //convert all extensions to lower case
            path.extname = path.extname.toLocaleLowerCase();
        }))
        .pipe(gulp.dest(dataTargetPath));
});

gulp.task('watch::copy-data', function () {

    const watch = require('gulp-watch');

    function fileAdded(event) {
        console.log('file added', event);
    }

    function fileRemoved(event) {
        console.log('file unlinked', event);
    }

    function fileChanged(event) {
        console.log('file changed', event);
    }

    return watch(dataSlug)
        .on('add', fileAdded).on('change', fileChanged).on('unlink', fileRemoved)
        .pipe(rename((path) => {
            //convert all extensions to lower case
            path.extname = path.extname.toLocaleLowerCase();
        }))
        .pipe(gulp.dest(dataTargetPath));
});


const rFileExt = /\.([0-9a-z]+)(?:[\?#]|$)/i;

function fileExtensionFromPath(path) {
    //get extension of file
    let fileExtMatch = path.match(rFileExt);
    if (fileExtMatch !== null) {
        let fileExt = fileExtMatch[fileExtMatch.length - 1];
        return fileExt;
    }
    return null;
}

let assetListFilePathStripped = "preloaderAssetList.json";
let assetListFilePath = dataSourcePath + "/" + assetListFilePathStripped;
let assetRootPath = "/assets";

/**
 *
 * @param {IncomingMessage} req
 * @param {ServerResponse} res
 * @param {function(err:*)} next
 */
function middleWareDisableCache(req, res, next) {
    res.setHeader('Cache-Control', 'private, no-cache, no-store, must-revalidate');
    res.setHeader('Expires', '-1');
    res.setHeader('Pragma', 'no-cache');
    next();
}

const releaseDestination = './release';

const middlewareReleaseCopier = (function () {
    const path = require('path');

    const records = {};

    function processRequest(req, res, next) {

        let url = req.url;

        //strip leading slashes
        while (url.charAt(0) === "/") {
            url = url.substr(1);
        }

        const decodedURL = decodeURIComponent(url);


        const dest = releaseDestination + '/' + decodedURL;

        const destPath = path.dirname(dest);


        try {
            fs.mkdirSync(destPath, { recursive: true });

            fs.copyFileSync('./public/' + decodedURL, dest, fs.constants.COPYFILE_FICLONE);

            if (decodedURL.trim().length > 0 && !records[decodedURL]) {
                //record
                records[decodedURL] = true;
                console.log('wrote: ', decodedURL);
            }
        } catch (e) {
            console.error("Failed to copy asset", e);
        }


        next();
    }

    return processRequest;
})();

const middleWareAssetListKeeper = (function () {


    function isAssetPath(path) {
        return path.indexOf(assetRootPath) === 0;
    }


    let assetHash = {};

    function guessAssetType(url, ext) {
        let assetDirectory = url.substring(assetRootPath.length);
        while (assetDirectory.charAt(0) === "/") {
            assetDirectory = assetDirectory.substr(1);
        }
        let iSlash = assetDirectory.indexOf("/");
        if (iSlash === -1) {
            assetDirectory = "";
        } else {
            assetDirectory = assetDirectory.substr(0, iSlash);
        }
        switch (ext) {
            case "json":
                switch (assetDirectory) {
                    case "models":
                        return "three.js";
                    case "levels":
                        return "level";
                    default:
                        return "json";
                }
            case "jpg":
            case "jpeg":
            case "png":
                return "image";
            case "ogg":
            case "mp3":
            //NOTE currently chrome doesn't seem to load these
            // return "sound";
            default :
                return null;
        }
    }

    function assetLevelByType(type) {
        switch (type) {
            case "image":
            case "three.js":
                return 1;
            case "level":
                return 0;
            case "sound":
            default :
                return 2;
        }
    }

    function tryRegisterAsset(url, ext) {
        if (!assetHash.hasOwnProperty(url)) {
            let type = guessAssetType(url, ext);
            if (type === null) {
                //ignore
                return;
            }
            let level = assetLevelByType(type);
            assetHash[url] = {
                "uri": url,
                "type": type,
                "level": level
            };
            writeAssetList();
        }
    }

    function writeAssetList() {
        let fileContents = [];
        for (let url in assetHash) {
            if (assetHash.hasOwnProperty(url)) {
                let urlStripped = url.substr(assetRootPath.length);
                if (urlStripped === assetListFilePathStripped) {
                    continue; //ignore file to which write will happen
                }
                fileContents.push(assetHash[url]);
            }
        }
        fs.writeFile(assetListFilePath, JSON.stringify(fileContents, 3, 3), function (err) {
            if (err) {
                return console.log(err);
            }
        });
    }

    function processRequest(req, res, next) {

        let url = req.url;
        if (isAssetPath(url)) {
            //strip leading slashes
            while (url.charAt(0) === "/") {
                url = url.substr(1);
            }
            let ext = fileExtensionFromPath(url);
            if (ext !== null) {
                tryRegisterAsset(url, ext)
            }
        }
        next();
    }

    return processRequest;
})();

gulp.task('clear-asset-list', function (done) {
    //clear out asset lists
    fs.writeFileSync(assetListFilePath, "[]");
    fs.writeFileSync(dataTargetPath + "/" + assetListFilePathStripped, "[]");
    done();
});

gulp.task('server-asset-recorder', gulp.series('clear-asset-list', function () {
    notifier.notify({
        title: "Server",
        message: "booted"
    });


    connect.server({
        root: 'public',
        middleware: function (connect, opt) {
            return [middleWareAssetListKeeper];
        },
        port: 8081,
        livereload: true
    });
}));

gulp.task('server-release-builder', function () {
    notifier.notify({
        title: "Server",
        message: "booted"
    });


    connect.server({
        root: 'public',
        middleware: function (connect, opt) {
            return [middlewareReleaseCopier];
        },
        port: 8081,
        livereload: true
    });
});

gulp.task('server', function () {
    notifier.notify({
        title: "Server",
        message: "booted"
    });

    return connect.server({
        root: 'public',
        port: 8080,
        livereload: true,
        middleware: function () {
            return [
                middleWareDisableCache
            ];
        }
    });
});


let watchSettings = { interval: 500, delay: 100 };

gulp.task('watch::styles', gulp.parallel(
    () => {
        return gulp.watch(sassGlob, watchSettings, gulp.parallel('css'));
    },
    () => {
        //css
        return gulp.watch(styleGlob, watchSettings, gulp.parallel('css'));
    }
));

gulp.task('watch',
    gulp.parallel(
        'watch::styles',
        'watch::copy-data',
        'watch::localization',
        () => {
            return makeWatchBundler('bundle.js', 'app/src/main.js', './public', "./app/src");
        }
    )
);


gulp.task('watch::dev', gulp.series(
    gulp.parallel(
        'copy-data',
        'html',
        'css',
        'localization'
    ),
    gulp.parallel('watch::styles', 'watch::copy-data', 'watch::localization')
));

gulp.task('default', gulp.series(
    gulp.parallel(
        'copy-data',
        'html',
        'css',
        'localization'
    ),
    gulp.parallel('server', 'watch')
));
