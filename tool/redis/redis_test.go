package redis

import (
	"fmt"
	"testing"
)

//[redis]
//Addresses = "127.0.0.1:6379"
//Password = ""
//MaxIdle = 30
//MaxActive = 30
//IdleTimeout = 200
//Cluster = 0
//DBNumber = 7

func TestRedis(t *testing.T) {
	ins := &Instance{
		DBNumber:    0,
		Addresses:   []string{"47.115.204.44:6379"},
		Password:    "",
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 200,
	}
	redisEngine, err := InitEngine(ins)
	if err != nil {
		t.Fatal(err)
	}
	//err = redisEngine.Set("test1", "hello world", time.Minute*1)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = redisEngine.Set("test2", "hello world2", 0)
	//if err != nil {
	//	t.Fatal(err)
	//}
	value := redisEngine.Get("test1")
	fmt.Println(value)
}
