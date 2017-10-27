 function execute(someFunction,value){
	someFunction(value);
}

//匿名函数
execute(function(word){console.log(word)}, "hello")
