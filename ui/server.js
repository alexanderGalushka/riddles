// let express = require('express');
// let app = express();
// let path = require('path');
//
// const dir = "/Users/agalushka/go/src/github.com/alexanderGalushka/riddles/ui";
//
// app.get('/', function(req, res) {
//   res.sendFile(path.join(dir + '/index.html'));
// });
//
// app.get('/liquidFillGauge.js', function(req, res) {
//   res.sendFile(path.join(dir + '/liquidFillGauge.js'));
// });
//
// app.get('/solveRiddle.js', function(req, res) {
//   res.sendFile(path.join(dir + '/solveRiddle.js'));
//   res.sendFile(path.join(dir + '/waterGuages.js'));
// });
//
// app.get('/waterGuages.js', function(req, res) {
//   res.sendFile(path.join(dir + '/waterGuages.js'));
// });
//
// app.listen(8080);

var express = require('express');
var app = express();
var path = require('path');

app.use(express.static(path.join(__dirname)));
app.use("/styles", express.static(__dirname));
app.use("/images", express.static(__dirname + '/images'));
app.use("/scripts", express.static(__dirname + '/scripts'));

// viewed at based directory http://localhost:8080/
app.get('/', function (req, res) {
  res.sendFile(path.join(__dirname + '/views/index.html'));
});

// add other routes below
app.get('/about', function (req, res) {
  res.sendFile(path.join(__dirname + 'views/about.html'));
});

app.listen(process.env.PORT || 8080);