var EventEmitter = require('events').EventEmitter;

var event = new EventEmitter;
event.on('some_event', function(){
	console.log("处理函数");
});

setTimeout(function(){
	event.emit('some_event');
},1000);