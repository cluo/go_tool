package main

import (
	"fmt"
	"gopkg.in/redis.v4"
)



func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "smart2016", // no password set
		DB:       0,           // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}

func main() {

}
