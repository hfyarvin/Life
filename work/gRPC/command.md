`
protoc --proto_path=IMPORT_PATH 
	--cpp_out=DST_DIR 
	--java_out=DST_DIR 
	--python_out=DST_DIR 
	--go_out=DST_DIR 
	--ruby_out=DST_DIR 
	--javanano_out=DST_DIR 
	--objc_out=DST_DIR 
	--csharp_out=DST_DIR 
	path/to/file.proto
`
params:
	-I： 不指定时默认查找当前目录,导入文件
	--go_out:
		- plugins=plugin1+plugin2 (plugins=grpc)
		- M 指定导入的.proto文件路径编译后对应的golang包名
		- import_prefix=xxx 为所有import路径添加前缀
		- import_path=foo/bar 用于指定未声明package或go_package的文件的包名
		- 末尾 :编译文件路径 .proto文件路径(支持通配符)