var hello = require('./hello');
var p = require('./person');

hello.world();

person = new p();
person.setName('arvin');
person.sayHello();

