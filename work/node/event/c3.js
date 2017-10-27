var events = require('events');
var eventEmitter = new events.EventEmitter();

var listener1 = function listener1() {
	console.log('This is listener1...');
}

var listener2 = function listener2() {
	console.log('This is listener2...');
}

eventEmitter.addListener('connection', listener1);

// 绑定 connection 事件，处理函数为 listener2
eventEmitter.on('connection', listener2);

var eventListeners = require('events').EventEmitter.listenerCount(eventEmitter,'connection');
console.log(eventListeners + " 个监听器监听连接事件。");

eventEmitter.emit('connection');
eventEmitter.removeListener('connection', listener1);

console.log("listener1 不再受监听。");

eventEmitter.emit('connection');

eventListeners = require('events').EventEmitter.listenerCount(eventEmitter,'connection');

console.log(eventListeners + '个监听器监听连接事件。');

console.log('程序执行完毕。');