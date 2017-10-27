var server = require("./server");
var router = require("./router");
console.log(__filename);
server.start(router.route);