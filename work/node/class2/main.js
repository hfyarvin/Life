var events = require('events');

var eventEmitter = new events.EventEmitter();

var connectHandler = function connected() {
	console.log('连接成功。');

	//触发
	eventEmitter.emit('data_received');
} //把该connected函数作为值传给connectHandler

eventEmitter.on('connection', connectHandler);

eventEmitter.on('data_received',function(){
	console.log('数据接受成功。');
});

console.log("开始触发connection...");

eventEmitter.emit('connection');

console.log("程序执行完毕。");