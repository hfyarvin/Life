### Key
- del k
- dump k 
- exists k
- expire k
- expireat k timestamp

### string
- set k v
- get k  
- getrange key start end 范围
- getset k new_v
- mget k1 k2 k3...
- append k v
- decr k 减一

### Hash
- hmset new_hash field1 v1 field2 v2
- hgetall new_hash
- hdel new_hash field1 field2...
- hget new_hash fieldn
- hkeys new_hash 所有自断
- hlen new_hash

### List
- lpush new_list v1
- lrange new_list 0 10

### Set(无序)
- sadd new_set k1
- smembers new_set

### Zset(有序)
- zadd new_zset 0 k1
- zrangebyscore new_set 0 100

### 其他
- redis-cli
- redis-cli -h hosh -p port -a password