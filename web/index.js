var express = require('express');
var app = express();
app.use('/static', express.static('public'));
app.get('/', function (req, res) {
	  res.redirect("/static/index.html")
});

app.listen(3000, function () {
	  console.log('Example app listening on port 3000!');
});
