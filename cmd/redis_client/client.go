package main

import (
	"context"
	"flag"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the listened address")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: *addr,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to ping-pong: %v", err)
	}

	_, err = client.Do(ctx,"bmadd", "test1", 1).Result() // int64
	if err != nil {
		log.Fatalf("failed to bmadd: %v", err)
	}

	_, err = client.Do(ctx,"bmaddmany", "test1", 2, 3, 10, 11).Result() // int64
	if err != nil {
		log.Fatalf("failed to bmaddmany: %v", err)
	}

	//_, err = client.Do(ctx,"bmaddmany", "test2", 1, 2, 3, 20, 21, 100).Result() // int64
	//if err != nil {
	//	log.Fatalf("failed to bmaddmany: %v", err)
	//}

	_, err = client.Do(ctx,"bmdiffstore", "test3", "test1", "test2").Result() // int64
	if err != nil {
		log.Fatalf("failed to bmaddmany: %v", err)
	}

	res, err := client.Do(ctx,"bmexists", "test2", 100).Result() // int64
	println(res.(int64))
	if err != nil {
		log.Fatalf("failed to bmaddmany: %v", err)
	}
	if res.(int64) != 1 {
		log.Fatalf("expect exists but found none (0)")
	}

	res, err = client.Do(ctx,"bmexists", "test3", 20).Result() // int64
	if err != nil {
		log.Fatalf("failed to bmaddmany: %v", err)
	}
	if res.(int64) != 0 {
		log.Fatalf("expect not found but found one (1)")
	}

	//client.Do(ctx, "bmsave")
}
