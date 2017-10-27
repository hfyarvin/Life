var fs = require("fs");

var data = '';

var readerStream = fs.createReadStream('input.txt');

readerStream.setEncoding('UTF8');

readerStream.on('data', function(chunk){
	console.log("process data");
	data += chunk;
});

readerStream.on('end',function(){
	console.log("process end");
	console.log(data);
});

readerStream.on('error', function(err){
	console.log("process error");
	console.log(err.stack);
});

console.log("The program is over.");

