var fs = require("fs");

var data = 'https://github.com/hfyarvin/Life'
// 创建写入流
var writeStream = fs.createWriteStream('output.txt');
writeStream.write(data,'UTF8');
writeStream.end();//标记文件末尾

writeStream.on('finish', function(){
	console.log("process finish");
	console.log("complete finish");
});
writeStream.on('error', function(err){
	console.log("process error");
   console.log(err.stack);
});

console.log("The program is over.");
