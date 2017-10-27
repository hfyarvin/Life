### redis-server.exe redis.windows.conf
### redis-cli.exe -h 127.0.0.1 -p 6379  
### config get * /config get loglevel / config set loglevel 'notice'

### set a bb  --> get a --> del a -->exists a
- getrange key start end  子字符串
- getset key value 返回旧值
- decr key 减一
- decr key decrement 减指定值
- append key value 追加字符串

### PING 
out> PONG

### move key db

### expire key
### expireat key timestamp

### rename key newkey

### type key