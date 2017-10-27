//阻塞
var fs = require("fs"); 
var data = fs.readFileSync('input.txt'); //同文件夹

console.log(data.toString());
console.log("Program is over!")