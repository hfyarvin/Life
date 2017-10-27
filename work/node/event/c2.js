//多个监听事件, 依次执行
var EventEmitter = require('events').EventEmitter;

var event = new EventEmitter;
event.on('some_event', function(arg1, arg2){
	console.log("处理函数", arg1, arg2);
});

event.on('some_event', function(arg1, arg2){
	console.log("处理函数2", arg1, arg2);
});

setTimeout(function(){
	event.emit('some_event', 'a', 'b');
},1000);