buf = new Buffer(26);

for (var i = 0; i < buf.length; i++) {
	buf[i] = i + 97;
}

console.log( buf.toString('ascii'));
console.log( buf.toString('ascii', 0, 5));
console.log( buf.toString('utf8', 0, 5));
console.log( buf.toString(undefined, 0, 5));

console.log(buf.toJSON(buf));