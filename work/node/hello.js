console.info("程序开始执行：");
console.log(__filename);
console.log(__dirname);
console.log("Hello World");

function printHello(word){
	console.log(word);
}

//延时函数
var t = setTimeout(printHello, 2000,"Hello this") //后面接参数
clearTimeout(t);//清除延时函数

//
// var t2 = setInterval(printHello, 2000, "lalala")
console.info("程序执行完毕。")
