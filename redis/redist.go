package redis

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strings"
	"time"
)

// TryCommand 试用相关命令.
func TryCommand()  {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(
				"redis://127.0.0.1:6379",
			)
		},
	}
	ctx := context.Background()
	conn, err := pool.GetContext(ctx)
	if err != nil {
		panic(err)
	}

	// test EXISTS命令.
	// 输出：false  <nil>.
	resp, err := redis.Bool(conn.Do("EXISTS", "h2"))
	fmt.Println(resp, err)

	// test hget一个不存在的key.
	// 输出：0 true.
	resp2, err := redis.Int64(conn.Do("HGET", "h2", "f1"))
	fmt.Println(resp2, err == redis.ErrNil)

	// test hget一个存在的key && 不存在的field.
	// 输出：0 true.
	resp3, err := redis.Int64(conn.Do("HGET", "h1", "dafsa"))
	fmt.Println(resp3, err == redis.ErrNil)

	// test hget一个存在的key && 存在的field.
	// 输出：100 <nil>.
	resp4, err := redis.Int64(conn.Do("HGET", "h1", "f3"))
	fmt.Println(resp4, err)

	// test hset一个不存在的field时：查看数据reply会是什么.
	// 输出：true <nil>.
	resp5, err := redis.Bool(conn.Do("HSET", "h1", "f5", 300))
	fmt.Println(resp5, err)

	// test hset一个存在的field时 && 会覆盖当前值.
	// 输出：false <nil>.
	resp6, err := redis.Bool(conn.Do("HSET", "h1", "f5", 400))
	fmt.Println(resp6, err)


	args := make([]interface{}, len([]string{"f1", "f7", "f2"})+1)
	args[0] = "h1"
	for i, key := range []string{"f1", "f7", "f2"} {
		args[i+1] = key
	}
	resp7, err := redis.Int64s(conn.Do("HMGET", args...))
	fmt.Println(resp7, err)

	resp8, err := redis.Strings(conn.Do("HMGET", "h1", "f4", "f5", "f6"))
	for k, r := range resp8 {
		fmt.Println(k, r=="")
	}
	fmt.Println(resp8, err)

	resp9, err := redis.Bool(conn.Do("HDEL", "h2", "f1"))
	fmt.Printf("--------resp9--------")
	fmt.Println(resp9, err)

	// 测试hscan命令
	var cursor int64
	hVals := make(map[string]string, 1000)
	for {
		vals, err := redis.Values(conn.Do("HSCAN", "h1", cursor, "COUNT", 1000))
		if err != nil {
			fmt.Printf("hscan err: %v\n", err)
			return
		}

		newCursor, err := redis.Int64(vals[0], nil)
		if err != nil {
			fmt.Printf("newCursor err: %v\n", err)
			return
		}

		cursor = newCursor
		hVals, err = redis.StringMap(vals[1], nil)
		if err != nil {
			fmt.Printf("StringMap err: %v\n", err)
			return
		}

		fmt.Println(hVals)

		if newCursor == 0 {
			fmt.Println("newCursor equal 0")
			break
		}
	}

	// 测试ZRANGEBYSCORE.
	zmembers, err := redis.StringMap(conn.Do("ZRANGEBYSCORE", "z1", 0, 2, "WITHSCORES"))
	fmt.Printf("ZRANGEBYSCORE resp: %+v, err: %+v", zmembers, err)

	// test zset.
	// test zadd  当member没在集合中，则插入会返回true, nil.
	zsetResp, err := redis.Bool(conn.Do("ZADD", "z1", "7", "m7"))
	fmt.Printf("zadd resp: %v, err: %v\n", zsetResp, err == nil)

	// test zrem.
	// 当该member在该集合中和返回 true, nil.
	zsetResp2, err := redis.Bool(conn.Do("ZREM", "z1", "m7"))
	fmt.Printf("zrem exist member resp: %v, err: %v\n", zsetResp2, err)

	// 当该member不在该集合中时，返回 false, nil.
	zsetResp3, err := redis.Bool(conn.Do("ZREM", "z1", "m100"))
	fmt.Printf("zrem not exist member resp: %v, err: %v\n", zsetResp3, err)

}

// InitUserIndustryIds 初始化用户行业ids数据.
func InitUserIndustryIds() {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(
				"redis://127.0.0.1:6379",
			)
		},
	}
	ctx := context.Background()
	conn, err := pool.GetContext(ctx)
	if err != nil {
		panic(err)
	}

	keyPrefix := "reportv2:idi"
	str := "100000001,100000002,100000003,100000004,100000005,100000006,100000007,100000008,100000009,100000010,100000011,100000012,100000013,100001001,100001002,100001003,100001004,100001005,100001006,100001007,100001008,100002001,100002002,100002003,100003001,100003002,100003003,100003004,100003005,100003006,100004001,100004002,100004003,100004004,100004005,100004006,100004007,100005001,100005002,100005003,100005004,100005005,100005006,100005007,100005008,100005009,100005010,100006001,100006002,100006003,100006004,100006005,100006006,100006007,100006008,100006009,100006010,100007001,100007002,100007003,100007004,100007005,100007006,100008001,100008002,100008003,100008004,100008005,100009001,100010001"
	oldIds := strings.Split(str, ",")
	timeNow := time.Now().Unix()
	for i := 1; i<=100; i++ {
		oldId1 := oldIds[rand.Intn(len(oldIds))]
		oldId2 := oldIds[rand.Intn(len(oldIds))]
		oldId3 := oldIds[rand.Intn(len(oldIds))]

		// 写入到zset中.
		ts1 := timeNow - int64(3600 * rand.Intn(len(oldIds)))
		ts2 := timeNow - int64(3600 * rand.Intn(len(oldIds)))
		ts3 := timeNow - int64(3600 * rand.Intn(len(oldIds)))

		key := fmt.Sprintf("%s:%d", keyPrefix, i)

		conn.Do("ZADD", key, ts1, oldId1)
		conn.Do("ZADD", key, ts2, oldId2)
		conn.Do("ZADD", key, ts3, oldId3)

		// 设置key有效期为一个月.
		conn.Do("EXPIRE", key, timeNow + 3600 * 24 * 30 - int64(3600 * i))
	}

	fmt.Println("init done")
}
