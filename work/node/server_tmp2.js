var http = require('http');//载入module

function onRequest(request, response){
	response.writeHead(200, {'Content-Type': 'text/plain'});

	response.end('Hello World\n');
}
http.createServer(onRequest).listen(8888);

console.log('Server running at http://127.0.0.1:8888');