package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
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

	startTime := gtime.Now()

	//for i := 0; i < 1000000; i++ {
	//	_, err = client.Do(ctx,"bmadd", "hdcj:lottery_bits:10000", i).Result() // int64
	//	if err != nil {
	//		log.Fatalf("failed to bmadd: %v", err)
	//	}
	//}

	for i := 0; i < 100000; i++ {
		res, err := client.Do(ctx,"bmexists", "hdcj:lottery_bits:10000", i).Result() // int64
		if err != nil {
			log.Fatalf("failed to bmadd: %v", err)
		}
		println(res.(int64))
	}

	endOpenTime := gtime.Now()

	takeOpenTime := endOpenTime.Sub(startTime).Minutes()

	println(fmt.Sprintf("%.2f", takeOpenTime))

	println("end")
}
