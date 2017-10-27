var http = require('http');
var fs = require('fs');
var url = require('url');

function onRequest(request, response){
	//解析
	var pathname = url.parse(request.url).pathname;

	console.log("Request for " + pathname + " received.");

	fs.readFile(pathname.substr(1), function(err,data){
		if (err) {
			console.log(err);
	         // HTTP 状态码: 404 : NOT FOUND
	         // Content Type: text/plain
	         response.writeHead(404, {'Content-Type': 'text/html'});
		}else{
	         // HTTP 状态码: 200 : OK
	         // Content Type: text/plain
	         response.writeHead(200, {'Content-Type': 'text/html'});    
	         
	         // 响应文件内容
	         response.write(data.toString());    
		}

		response.end();
	});
}

//创建服务器
http.createServer(onRequest).listen(8081);

console.log('Server running at http://127.0.0.1:8081/');