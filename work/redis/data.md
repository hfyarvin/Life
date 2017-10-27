### DATA_TYPE: string; hash; list; set; zset

## string
## hash
- HMSET runoobkey name "redis tutorial" description "redis basic commands for caching" likes 20 visitors 23000
- HGETALL runoobkey 
- hdel key field1 ...删键值对
- hexists key field
- hincrby key field increment 增值
- hincrbyfloat key field increment
- hkeys key
- hlen key
- hmget key field1 ...获取指定键值对
- hmset key f1 v1 f2 v2... 设置
- hset key field value 设值
- hsetnx key field value 只有当field不存在，才设定值
- hvals key 所有值

## list
- lpush runoob redis
- lpush runoob mongodb
- lpush runoob rabitmq
- lrange runoob 0 10 //倒序取，先进后出
- blpop key1 [key2] timeout
- llen key
- rpop key 移除并获取最后一个 
- rpoplpush source destination 移除最后一个并添加到另一个列表
- rpush key v1 v2 ...添加


## set //元素唯一性
- sadd runoob redis  
- sadd runoob mongo
- sadd key member1 member2...
- smembers runoob
- scard key 成员数
- sdiff key1 key2 ...给定所有集合的差值
- sdiff destination key1 key2 ... 差值存储在destination中

## zset
- zadd key score member 
- zadd runoob 0 mongodb
- ZRANGEBYSCORE runoob 0 1000

### HyperLogLog
- pfadd key element [element ...]
- pfcount key [key ...]
- pfmerge destkey sourcekey [sourcekey ...]

### 订阅
- subscribe redischat
- publish redischat --> publish channel message
- psubscribe pattern [pattern]
- pubsub subcommand[args...]
- punsubscribe pattern 退订
- subscribe channel
- unsubscribe channel

### 事务
- multi 开始
- exec 结束
- discard 放弃
- unwatch 放弃对key所欲监视
- watch key 监视

### 脚本
- eval script numkeys key [key ...] arg [arg ...]
- EVAL "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}" 2 key1 key2 first second

### 连接
- auth password <-- CONFIG SET requirepass "mypass"
- echo string
- quit
- select index
	set db_number 0
	select 1
	get db_number
	set db_number 1
	get db_number
	select 3

### 服务器
- info
- dbsize
- flushall