var fs = require("fs");

var readerStream = fs.createReadStream('input.txt');

var writeStream = fs.createWriteStream('output.txt');

readerStream.pipe(writeStream); //管道读写

console.log("The program is over.");