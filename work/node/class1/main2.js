//非阻塞，不按顺序执行
var fs = require("fs");

fs.readFile('input.txt', function (err, data) { //readFile:异步函数读文件
    if (err) return console.error(err);

    console.log(data.toString());
});

console.log("程序执行结束!");