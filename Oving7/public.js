var https = require('https');
var http = require('http');
var express = require('express');
var fs = require('fs');

var app = express();
app.use(express.static('public'));

var key = fs.readFileSync('sslcert/server.key', 'utf-8');
var cert = fs.readFileSync('sslcert/server.crt', 'utf-8');

var credentials = {key, cert};
var http_app = express();
http_app.get('*', function(req, res) {
    return res.redirect('https://' + req.headers.host + req.url)
});

var httpserver = http.createServer(http_app);
var httpsserver = https.createServer(credentials, app);
httpsserver.listen(443);
httpserver.listen(80);
console.log("listening");