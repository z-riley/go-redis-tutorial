package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // use default DB
	})

	ctx := context.Background() // create empty context

	// Test the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Store a string
	status := client.Set(ctx, "my key", "my value", 0)
	if status.Err() != nil {
		panic(status.Err())
	}

	// Retrieve a string
	val, err := client.Get(ctx, "my key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("my key:", val)

	// Store a map
	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	for key, val := range session {
		err := client.HSet(ctx, "my map", key, val).Err()
		if err != nil {
			panic(err)
		}
	}

	// Retrieve a map
	userSession := client.HGetAll(ctx, "my map").Val()
	fmt.Println(userSession)

}
