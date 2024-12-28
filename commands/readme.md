Key Commands:
DEL - Remove one or more keys
EXISTS - Check if a key exists
EXPIRE - Set a timeout on a key
TTL - Get the time-to-live (TTL) of a key
PERSIST - Remove the timeout on a key
TYPE - Get the type of a key
RENAME - Rename a key
RENAMENX - Rename a key, only if the new key does not exist
KEYS - Find all keys matching a pattern
SCAN - Incrementally iterate the keys space

String Commands:
SET - Set the value of a key
GET - Get the value of a key
APPEND - Append a value to a key
INCR - Increment the integer value of a key
DECR - Decrement the integer value of a key
MGET - Get the values of multiple keys
MSET - Set multiple keys to multiple values
SETNX - Set the value of a key only if it does not exist
SETEX - Set the value of a key with an expiration time
PSETEX - Set the value of a key with millisecond expiration time

List Commands:
LPUSH - Prepend one or more values to a list
RPUSH - Append one or more values to a list
LPOP - Remove and get the first element of a list
RPOP - Remove and get the last element of a list
LRANGE - Get a range of elements from a list
LLEN - Get the length of a list
LREM - Remove elements from a list
LSET - Set the value of an element in a list by its index
LINSERT - Insert an element before or after another element in a list

Set Commands:
SADD - Add one or more members to a set
SREM - Remove one or more members from a set
SMEMBERS - Get all members of a set
SISMEMBER - Check if a member exists in a set
SCARD - Get the number of members in a set
SMOVE - Move a member from one set to another
SINTER - Get the intersection of multiple sets
SUNION - Get the union of multiple sets
SDIFF - Get the difference between multiple sets

Sorted Set Commands (Score-related):
ZADD - Add one or more members to a sorted set
ZREM - Remove one or more members from a sorted set
ZSCORE - Get the score of a member in a sorted set
ZINCRBY - Increment the score of a member in a sorted set
ZRANGE - Get the members of a sorted set in a specific rank range
ZRANGEBYSCORE - Get the members of a sorted set in a specific score range
ZRANK - Get the rank of a member in a sorted set
ZCARD - Get the number of members in a sorted set
ZREMRANGEBYSCORE - Remove members from a sorted set by score
ZREVRANGE - Get the members of a sorted set in reverse order
ZREVRANK - Get the rank of a member in a sorted set (reverse order)

Hash Commands:
HSET - Set the value of a field in a hash
HGET - Get the value of a field in a hash
HGETALL - Get all fields and values in a hash
HDEL - Remove one or more fields from a hash
HEXISTS - Check if a field exists in a hash
HINCRBY - Increment the value of a field in a hash by a given number
HKEYS - Get all field names in a hash
HVALS - Get all values in a hash
HLEN - Get the number of fields in a hash
HMSET - Set multiple fields in a hash (deprecated in newer Redis versions, use HSET)

Pub/Sub Commands:
PUBLISH - Publish a message to a channel
SUBSCRIBE - Subscribe to one or more channels
UNSUBSCRIBE - Unsubscribe from one or more channels
PSUBSCRIBE - Subscribe to one or more patterns
PUNSUBSCRIBE - Unsubscribe from one or more patterns

Transaction Commands:
MULTI - Start a transaction
EXEC - Execute all commands in a transaction
DISCARD - Discard all commands in a transaction
WATCH - Watch one or more keys for changes

Scripting Commands:
EVAL - Execute a Lua script
EVALSHA - Execute a Lua script by its SHA1 hash
SCRIPT LOAD - Load a Lua script into Redis
SCRIPT EXISTS - Check if a Lua script exists in Redis
SCRIPT FLUSH - Flush all Lua scripts from Redis

Connection Commands:
AUTH - Authenticate to the server
PING - Ping the server
QUIT - Close the connection

Server Commands:
INFO - Get information and statistics about the server
CONFIG GET - Get the configuration parameters of the server
CONFIG SET - Set the configuration parameters of the server
FLUSHDB - Remove all keys from the current database
FLUSHALL - Remove all keys from all databases
SAVE - Synchronously save the dataset to disk
BGSAVE - Asynchronously save the dataset to disk
SHUTDOWN - Stop the server

HyperLogLog Commands:
PFADD - Add elements to a HyperLogLog
PFCOUNT - Count unique elements in a HyperLogLog
PFMERGE - Merge multiple HyperLogLogs into one

Geospatial Commands:
GEOADD - Add one or more geospatial elements to a geospatial index
GEODIST - Get the distance between two members of a geospatial index
GEORADIUS - Get all the members in a geospatial index within a radius
GEORADIUSBYMEMBER - Get all the members in a geospatial index within a radius of a member

Bitmap Commands:
SETBIT - Set the value of a bit in a string
GETBIT - Get the value of a bit in a string
BITCOUNT - Count the number of set bits in a string
BITOP - Perform bitwise operations on one or more strings
