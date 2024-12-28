package main

import (
	"log"
	"testing"
)

func TestProtocolReader(t *testing.T) {

	var msg string

	//// DEL - Remove one or more keys
	//msg += "*2\r\n$3\r\nDEL\r\n$5\r\nmykey\r\n"
	//
	//// EXISTS - Check if a key exists
	//msg += "*2\r\n$6\r\nEXISTS\r\n$5\r\nmykey\r\n"
	//
	//// EXPIRE - Set a timeout on a key (EXPIRE mykey 60)
	//msg += "*3\r\n$6\r\nEXPIRE\r\n$5\r\nmykey\r\n$2\r\n60\r\n"
	//
	//// TTL - Get the time-to-live (TTL) of a key
	//msg += "*2\r\n$3\r\nTTL\r\n$5\r\nmykey\r\n"
	//
	//// PERSIST - Remove the timeout on a key
	//msg += "*2\r\n$7\r\nPERSIST\r\n$5\r\nmykey\r\n"
	//
	//// TYPE - Get the type of a key
	//msg += "*2\r\n$4\r\nTYPE\r\n$5\r\nmykey\r\n"
	//
	//// RENAME - Rename a key
	//msg += "*3\r\n$6\r\nRENAME\r\n$5\r\nmykey\r\n$6\r\nnewkey\r\n"
	//
	//// RENAMENX - Rename a key, only if the new key does not exist
	//msg += "*3\r\n$8\r\nRENAMENX\r\n$5\r\nmykey\r\n$6\r\nnewkey\r\n"
	//
	//// KEYS - Find all keys matching a pattern
	//msg += "*2\r\n$4\r\nKEYS\r\n$8\r\npattern*\r\n"
	//
	//// SCAN - Incrementally iterate the keys space (SCAN 0 MATCH pattern*)
	//msg += "*3\r\n$4\r\nSCAN\r\n$1\r\n0\r\n$6\r\nMATCH\r\n$8\r\npattern*\r\n"

	// SET - Set the value of a key
	msg += "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nvalue\r\n"
	//msg += "*2\r\n$3\r\nSET\r\n$5\r\nmykey\r\n"

	//// GET - Get the value of a key
	//msg += "*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n"
	//
	//// APPEND - Append a value to a key
	//msg += "*3\r\n$6\r\nAPPEND\r\n$5\r\nmykey\r\n$6\r\nnewvalue\r\n"
	//
	//// INCR - Increment the integer value of a key
	//msg += "*2\r\n$4\r\nINCR\r\n$7\r\ncounter\r\n"
	//
	//// DECR - Decrement the integer value of a key
	//msg += "*2\r\n$4\r\nDECR\r\n$7\r\ncounter\r\n"
	//
	//// MGET - Get the values of multiple keys
	//msg += "*3\r\n$4\r\nMGET\r\n$4\r\nkey1\r\n$4\r\nkey2\r\n"
	//
	//// MSET - Set multiple keys to multiple values
	//msg += "*5\r\n$4\r\nMSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n$4\r\nkey2\r\n$6\r\nvalue2\r\n"
	//
	//// SETNX - Set the value of a key only if it does not exist
	//msg += "*4\r\n$5\r\nSETNX\r\n$5\r\nmykey\r\n$6\r\nvalue\r\n"
	//
	//// SETEX - Set the value of a key with an expiration time
	//msg += "*5\r\n$5\r\nSETEX\r\n$5\r\nmykey\r\n$6\r\nvalue\r\n$2\r\n60\r\n"
	//
	//// PSETEX - Set the value of a key with millisecond expiration time
	//msg += "*5\r\n$6\r\nPSETEX\r\n$5\r\nmykey\r\n$6\r\nvalue\r\n$5\r\n1000\r\n"
	//
	//// LPUSH - Prepend one or more values to a list
	//msg += "*3\r\n$5\r\nLPUSH\r\n$5\r\nmylist\r\n$6\r\nvalue1\r\n"
	//
	//// RPUSH - Append one or more values to a list
	//msg += "*3\r\n$5\r\nRPUSH\r\n$5\r\nmylist\r\n$6\r\nvalue1\r\n"
	//
	//// LPOP - Remove and get the first element of a list
	//msg += "*2\r\n$4\r\nLPOP\r\n$5\r\nmylist\r\n"
	//
	//// RPOP - Remove and get the last element of a list
	//msg += "*2\r\n$4\r\nRPOP\r\n$5\r\nmylist\r\n"
	//
	//// LRANGE - Get a range of elements from a list (LRANGE mylist 0 2)
	//msg += "*4\r\n$6\r\nLRANGE\r\n$5\r\nmylist\r\n$1\r\n0\r\n$1\r\n2\r\n"
	//
	//// LLEN - Get the length of a list
	//msg += "*2\r\n$4\r\nLLEN\r\n$5\r\nmylist\r\n"
	//
	//// LREM - Remove elements from a list (LREM mylist 0 value1)
	//msg += "*4\r\n$5\r\nLREM\r\n$5\r\nmylist\r\n$1\r\n0\r\n$6\r\nvalue1\r\n"
	//
	//// LSET - Set the value of an element in a list by its index (LSET mylist 0 new_value)
	//msg += "*4\r\n$4\r\nLSET\r\n$5\r\nmylist\r\n$1\r\n0\r\n$9\r\nnew_value\r\n"
	//
	//// LINSERT - Insert an element before or after another element in a list (LINSERT mylist BEFORE value1 new_value)
	//msg += "*5\r\n$7\r\nLINSERT\r\n$5\r\nmylist\r\n$7\r\nBEFORE\r\n$6\r\nvalue1\r\n$9\r\nnew_value\r\n"
	//
	//// SADD - Add one or more members to a set
	//msg += "*3\r\n$4\r\nSADD\r\n$5\r\nmyset\r\n$6\r\nmember1\r\n"
	//
	//// SREM - Remove one or more members from a set
	//msg += "*3\r\n$4\r\nSREM\r\n$5\r\nmyset\r\n$6\r\nmember1\r\n"
	//
	//// SMEMBERS - Get all members of a set
	//msg += "*2\r\n$8\r\nSMEMBERS\r\n$5\r\nmyset\r\n"
	//
	//// SISMEMBER - Check if a member exists in a set
	//msg += "*3\r\n$9\r\nSISMEMBER\r\n$5\r\nmyset\r\n$7\r\nmember1\r\n"
	//
	//// SCARD - Get the number of members in a set
	//msg += "*2\r\n$5\r\nSCARD\r\n$5\r\nmyset\r\n"
	//
	//// SMOVE - Move a member from one set to another
	//msg += "*4\r\n$5\r\nSMOVE\r\n$5\r\nset1\r\n$5\r\nset2\r\n$7\r\nmember1\r\n"
	//
	//// SINTER - Get the intersection of multiple sets
	//msg += "*3\r\n$6\r\nSINTER\r\n$5\r\nset1\r\n$5\r\nset2\r\n"
	//
	//// SUNION - Get the union of multiple sets
	//msg += "*3\r\n$6\r\nSUNION\r\n$5\r\nset1\r\n$5\r\nset2\r\n"
	//
	//// SDIFF - Get the difference between multiple sets
	//msg += "*3\r\n$6\r\nSDIFF\r\n$5\r\nset1\r\n$5\r\nset2\r\n"
	//
	//// ZADD - Add one or more members to a sorted set
	//msg += "*4\r\n$4\r\nZADD\r\n$5\r\nmyzset\r\n$2\r\n1\r\n$6\r\nmember1\r\n"
	//
	//// ZREM - Remove one or more members from a sorted set
	//msg += "*3\r\n$4\r\nZREM\r\n$5\r\nmyzset\r\n$6\r\nmember1\r\n"
	//
	//// ZSCORE - Get the score of a member in a sorted set
	//msg += "*3\r\n$6\r\nZSCORE\r\n$5\r\nmyzset\r\n$6\r\nmember1\r\n"
	//
	//// ZINCRBY - Increment the score of a member in a sorted set
	//msg += "*4\r\n$7\r\nZINCRBY\r\n$5\r\nmyzset\r\n$2\r\n2\r\n$6\r\nmember1\r\n"
	//
	//// ZRANGE - Get the members of a sorted set in a specific rank range
	//msg += "*4\r\n$6\r\nZRANGE\r\n$5\r\nmyzset\r\n$1\r\n0\r\n$1\r\n2\r\n"
	//
	//// ZRANGEBYSCORE - Get the members of a sorted set in a specific score range
	//msg += "*5\r\n$10\r\nZRANGEBYSCORE\r\n$5\r\nmyzset\r\n$2\r\n0\r\n$2\r\n10\r\n"
	//
	//// ZRANK - Get the rank of a member in a sorted set
	//msg += "*3\r\n$6\r\nZRANK\r\n$5\r\nmyzset\r\n$6\r\nmember1\r\n"
	//
	//// ZCARD - Get the number of members in a sorted set
	//msg += "*2\r\n$5\r\nZCARD\r\n$5\r\nmyzset\r\n"
	//
	//// ZREMRANGEBYSCORE - Remove members from a sorted set by score
	//msg += "*4\r\n$14\r\nZREMRANGEBYSCORE\r\n$5\r\nmyzset\r\n$2\r\n0\r\n$2\r\n10\r\n"
	//
	//// ZREVRANGE - Get the members of a sorted set in reverse order
	//msg += "*4\r\n$9\r\nZREVRANGE\r\n$5\r\nmyzset\r\n$1\r\n0\r\n$1\r\n2\r\n"
	//
	//// ZREVRANK - Get the rank of a member in a sorted set (reverse order)
	//msg += "*3\r\n$8\r\nZREVRANK\r\n$5\r\nmyzset\r\n$6\r\nmember1\r\n"
	//
	//// HSET - Set the value of a field in a hash
	//msg += "*4\r\n$4\r\nHSET\r\n$3\r\nmyhash\r\n$5\r\nfield\r\n$6\r\nvalue\r\n"
	//
	//// HGET - Get the value of a field in a hash
	//msg += "*3\r\n$4\r\nHGET\r\n$3\r\nmyhash\r\n$5\r\nfield\r\n"
	//
	//// HGETALL - Get all fields and values in a hash
	//msg += "*2\r\n$7\r\nHGETALL\r\n$3\r\nmyhash\r\n"
	//
	//// HDEL - Remove one or more fields from a hash
	//msg += "*3\r\n$4\r\nHDEL\r\n$3\r\nmyhash\r\n$5\r\nfield\r\n"
	//
	//// HEXISTS - Check if a field exists in a hash
	//msg += "*3\r\n$7\r\nHEXISTS\r\n$3\r\nmyhash\r\n$5\r\nfield\r\n"
	//
	//// HINCRBY - Increment the value of a field in a hash by a given number
	//msg += "*4\r\n$7\r\nHINCRBY\r\n$3\r\nmyhash\r\n$5\r\nfield\r\n$2\r\n10\r\n"
	//
	//// HKEYS - Get all field names in a hash
	//msg += "*2\r\n$5\r\nHKEYS\r\n$3\r\nmyhash\r\n"
	//
	//// HVALS - Get all values in a hash
	//msg += "*2\r\n$5\r\nHVALS\r\n$3\r\nmyhash\r\n"
	//
	//// HLEN - Get the number of fields in a hash
	//msg += "*2\r\n$4\r\nHLEN\r\n$3\r\nmyhash\r\n"
	//
	//// HMSET - Set multiple fields in a hash (deprecated, use HSET)
	//msg += "*4\r\n$5\r\nHMSET\r\n$3\r\nmyhash\r\n$5\r\nfield1\r\n$6\r\nvalue1\r\n$5\r\nfield2\r\n$6\r\nvalue2\r\n"
	//
	//// PUBLISH - Publish a message to a channel
	//msg += "*3\r\n$7\r\nPUBLISH\r\n$9\r\nmychannel\r\n$5\r\nHello\r\n"
	//
	//// SUBSCRIBE - Subscribe to one or more channels
	//msg += "*2\r\n$9\r\nSUBSCRIBE\r\n$5\r\nmychan\r\n"
	//
	//// UNSUBSCRIBE - Unsubscribe from one or more channels
	//msg += "*2\r\n$11\r\nUNSUBSCRIBE\r\n$5\r\nmychan\r\n"
	//
	//// PSUBSCRIBE - Subscribe to one or more patterns
	//msg += "*2\r\n$10\r\nPSUBSCRIBE\r\n$6\r\n*pattern\r\n"
	//
	//// PUNSUBSCRIBE - Unsubscribe from one or more patterns
	//msg += "*2\r\n$12\r\nPUNSUBSCRIBE\r\n$6\r\n*pattern\r\n"
	//
	//// MULTI - Start a transaction
	//msg += "*1\r\n$5\r\nMULTI\r\n"
	//
	//// EXEC - Execute all commands in a transaction
	//msg += "*1\r\n$4\r\nEXEC\r\n"
	//
	//// DISCARD - Discard all commands in a transaction
	//msg += "*1\r\n$7\r\nDISCARD\r\n"
	//
	//// WATCH - Watch one or more keys for changes
	//msg += "*2\r\n$5\r\nWATCH\r\n$5\r\nmykey\r\n"
	//
	//// EVAL - Execute a Lua script
	//msg += "*4\r\n$4\r\nEVAL\r\n$13\r\nreturn redis.call('set', 'foo', 'bar')\r\n$0\r\n\r\n"
	//
	//// EVALSHA - Execute a Lua script by its SHA1 hash
	//msg += "*3\r\n$7\r\nEVALSHA\r\n$40\r\n2bf48c2d0b3469ab7266b7c5c199f10a01f85c24\r\n$1\r\nfoo\r\n"
	//
	//// SCRIPT LOAD - Load a Lua script into Redis
	//msg += "*2\r\n$10\r\nSCRIPT\r\n$4\r\nLOAD\r\n"
	//
	//// SCRIPT EXISTS - Check if a Lua script exists in Redis
	//msg += "*2\r\n$13\r\nSCRIPT\r\n$7\r\nEXISTS\r\n$40\r\n2bf48c2d0b3469ab7266b7c5c199f10a01f85c24\r\n"
	//
	//// SCRIPT FLUSH - Flush all Lua scripts from Redis
	//msg += "*2\r\n$13\r\nSCRIPT\r\n$5\r\nFLUSH\r\n"
	//
	//// AUTH - Authenticate to the server
	//msg += "*2\r\n$4\r\nAUTH\r\n$6\r\npassword\r\n"
	//
	//// PING - Ping the server
	//msg += "*1\r\n$4\r\nPING\r\n"
	//
	//// QUIT - Close the connection
	//msg += "*1\r\n$4\r\nQUIT\r\n"
	//
	//// INFO - Get information and statistics about the server
	//msg += "*1\r\n$4\r\nINFO\r\n"
	//
	//// CONFIG GET - Get the configuration parameters of the server
	//msg += "*2\r\n$6\r\nCONFIG\r\n$3\r\nGET\r\n"
	//
	//// CONFIG SET - Set the configuration parameters of the server
	//msg += "*3\r\n$6\r\nCONFIG\r\n$3\r\nSET\r\n$10\r\nmaxmemory\r\n$5\r\n100mb\r\n"
	//
	//// FLUSHDB - Remove all keys from the current database
	//msg += "*1\r\n$7\r\nFLUSHDB\r\n"
	//
	//// FLUSHALL - Remove all keys from all databases
	//msg += "*1\r\n$8\r\nFLUSHALL\r\n"
	//
	//// SAVE - Synchronously save the dataset to disk
	//msg += "*1\r\n$4\r\nSAVE\r\n"
	//
	//// BGSAVE - Asynchronously save the dataset to disk
	//msg += "*1\r\n$7\r\nBGSAVE\r\n"
	//
	//// SHUTDOWN - Stop the server
	//msg += "*1\r\n$8\r\nSHUTDOWN\r\n"
	//
	//// PFADD - Add elements to a HyperLogLog
	//msg += "*2\r\n$5\r\nPFADD\r\n$6\r\nmyhyperlog\r\n$6\r\nelement1\r\n"
	//
	//// PFCOUNT - Count unique elements in a HyperLogLog
	//msg += "*2\r\n$7\r\nPFCOUNT\r\n$6\r\nmyhyperlog\r\n"
	//
	//// PFMERGE - Merge multiple HyperLogLogs into one
	//msg += "*3\r\n$7\r\nPFMERGE\r\n$6\r\nmymergedlog\r\n$6\r\nlog1\r\n$6\r\nlog2\r\n"
	//
	//// GEOADD - Add one or more geospatial elements to a geospatial index
	//msg += "*5\r\n$6\r\nGEOADD\r\n$9\r\nmygeospot\r\n$5\r\n13.3611\r\n$5\r\n52.2053\r\n$6\r\nPlace1\r\n"
	//
	//// GEODIST - Get the distance between two members of a geospatial index
	//msg += "*3\r\n$7\r\nGEODIST\r\n$9\r\nmygeospot\r\n$6\r\nPlace1\r\n$6\r\nPlace2\r\n"
	//
	//// GEORADIUS - Get all the members in a geospatial index within a radius
	//msg += "*5\r\n$9\r\nGEORADIUS\r\n$9\r\nmygeospot\r\n$5\r\n13.3611\r\n$5\r\n52.2053\r\n$3\r\n200\r\n$2\r\nkm\r\n"
	//
	//// GEORADIUSBYMEMBER - Get all the members in a geospatial index within a radius of a member
	//msg += "*4\r\n$13\r\nGEORADIUSBYMEMBER\r\n$9\r\nmygeospot\r\n$6\r\nPlace1\r\n$3\r\n200\r\n$2\r\nkm\r\n"
	//
	//// SETBIT - Set the value of a bit in a string
	//msg += "*4\r\n$6\r\nSETBIT\r\n$5\r\nmybitmap\r\n$3\r\n10\r\n$1\r\n1\r\n"
	//
	//// GETBIT - Get the value of a bit in a string
	//msg += "*3\r\n$6\r\nGETBIT\r\n$5\r\nmybitmap\r\n$3\r\n10\r\n"
	//
	//// BITCOUNT - Count the number of set bits in a string
	//msg += "*2\r\n$8\r\nBITCOUNT\r\n$5\r\nmybitmap\r\n"
	//
	//// BITOP - Perform bitwise operations on one or more strings
	//msg += "*4\r\n$5\r\nBITOP\r\n$3\r\nAND\r\n$5\r\nresult\r\n$5\r\nmybitmap\r\n"

	commands, err := ParseCommand(msg)
	if err != nil {
		panic(err)
	}
	log.Println(commands)
}

//Key Commands:
//DEL - Remove one or more keys
//EXISTS - Check if a key exists
//EXPIRE - Set a timeout on a key
//TTL - Get the time-to-live (TTL) of a key
//PERSIST - Remove the timeout on a key
//TYPE - Get the type of a key
//RENAME - Rename a key
//RENAMENX - Rename a key, only if the new key does not exist
//KEYS - Find all keys matching a pattern
//SCAN - Incrementally iterate the keys space
//
//String Commands:
//SET - Set the value of a key
//GET - Get the value of a key
//APPEND - Append a value to a key
//INCR - Increment the integer value of a key
//DECR - Decrement the integer value of a key
//MGET - Get the values of multiple keys
//MSET - Set multiple keys to multiple values
//SETNX - Set the value of a key only if it does not exist
//SETEX - Set the value of a key with an expiration time
//PSETEX - Set the value of a key with millisecond expiration time
//
//List Commands:
//LPUSH - Prepend one or more values to a list
//RPUSH - Append one or more values to a list
//LPOP - Remove and get the first element of a list
//RPOP - Remove and get the last element of a list
//LRANGE - Get a range of elements from a list
//LLEN - Get the length of a list
//LREM - Remove elements from a list
//LSET - Set the value of an element in a list by its index
//LINSERT - Insert an element before or after another element in a list
//
//Set Commands:
//SADD - Add one or more members to a set
//SREM - Remove one or more members from a set
//SMEMBERS - Get all members of a set
//SISMEMBER - Check if a member exists in a set
//SCARD - Get the number of members in a set
//SMOVE - Move a member from one set to another
//SINTER - Get the intersection of multiple sets
//SUNION - Get the union of multiple sets
//SDIFF - Get the difference between multiple sets
//
//Sorted Set Commands (Score-related):
//ZADD - Add one or more members to a sorted set
//ZREM - Remove one or more members from a sorted set
//ZSCORE - Get the score of a member in a sorted set
//ZINCRBY - Increment the score of a member in a sorted set
//ZRANGE - Get the members of a sorted set in a specific rank range
//ZRANGEBYSCORE - Get the members of a sorted set in a specific score range
//ZRANK - Get the rank of a member in a sorted set
//ZCARD - Get the number of members in a sorted set
//ZREMRANGEBYSCORE - Remove members from a sorted set by score
//ZREVRANGE - Get the members of a sorted set in reverse order
//ZREVRANK - Get the rank of a member in a sorted set (reverse order)
//
//Hash Commands:
//HSET - Set the value of a field in a hash
//HGET - Get the value of a field in a hash
//HGETALL - Get all fields and values in a hash
//HDEL - Remove one or more fields from a hash
//HEXISTS - Check if a field exists in a hash
//HINCRBY - Increment the value of a field in a hash by a given number
//HKEYS - Get all field names in a hash
//HVALS - Get all values in a hash
//HLEN - Get the number of fields in a hash
//HMSET - Set multiple fields in a hash (deprecated in newer Redis versions, use HSET)
//
//Pub/Sub Commands:
//PUBLISH - Publish a message to a channel
//SUBSCRIBE - Subscribe to one or more channels
//UNSUBSCRIBE - Unsubscribe from one or more channels
//PSUBSCRIBE - Subscribe to one or more patterns
//PUNSUBSCRIBE - Unsubscribe from one or more patterns
//
//Transaction Commands:
//MULTI - Start a transaction
//EXEC - Execute all commands in a transaction
//DISCARD - Discard all commands in a transaction
//WATCH - Watch one or more keys for changes
//
//Scripting Commands:
//EVAL - Execute a Lua script
//EVALSHA - Execute a Lua script by its SHA1 hash
//SCRIPT LOAD - Load a Lua script into Redis
//SCRIPT EXISTS - Check if a Lua script exists in Redis
//SCRIPT FLUSH - Flush all Lua scripts from Redis
//
//Connection Commands:
//AUTH - Authenticate to the server
//PING - Ping the server
//QUIT - Close the connection
//
//Server Commands:
//INFO - Get information and statistics about the server
//CONFIG GET - Get the configuration parameters of the server
//CONFIG SET - Set the configuration parameters of the server
//FLUSHDB - Remove all keys from the current database
//FLUSHALL - Remove all keys from all databases
//SAVE - Synchronously save the dataset to disk
//BGSAVE - Asynchronously save the dataset to disk
//SHUTDOWN - Stop the server
//
//HyperLogLog Commands:
//PFADD - Add elements to a HyperLogLog
//PFCOUNT - Count unique elements in a HyperLogLog
//PFMERGE - Merge multiple HyperLogLogs into one
//
//Geospatial Commands:
//GEOADD - Add one or more geospatial elements to a geospatial index
//GEODIST - Get the distance between two members of a geospatial index
//GEORADIUS - Get all the members in a geospatial index within a radius
//GEORADIUSBYMEMBER - Get all the members in a geospatial index within a radius of a member
//
//Bitmap Commands:
//SETBIT - Set the value of a bit in a string
//GETBIT - Get the value of a bit in a string
//BITCOUNT - Count the number of set bits in a string
//BITOP - Perform bitwise operations on one or more strings
