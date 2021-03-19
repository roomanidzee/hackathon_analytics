package main

import (
    "runtime"
    "fmt"
    "time"
    "context"
    "strings"
    "strconv"
    "github.com/go-redis/redis/v8"
)

func bRedisToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func insertRedisData(itemsNum int) {

	var ctx = context.Background()

        rdb := redis.NewClient(&redis.Options{
           Addr:     "localhost:6379",
           Password: "", // no password set
           DB:       0,  // use default DB
       })

	start := time.Now()
	for i := 0; i < itemsNum; i++ {
		item := newItem(i)

                var sb1 strings.Builder
                var sb2 strings.Builder

		sb1.WriteString(strconv.Itoa(int(item.ID)))

		sb2.WriteString(item.Name)
		sb2.WriteString("_")
		sb2.WriteString(item.Description)
		sb2.WriteString("_")
		sb2.WriteString(strconv.Itoa(item.Year))


		err := rdb.Set(ctx, sb1.String(), sb2.String(), 0).Err()
                if err != nil {
                      panic(err)
                }

	}

	var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc = %v MiB", bRedisToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bRedisToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bRedisToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)

	fmt.Println(time.Since(start))

}
