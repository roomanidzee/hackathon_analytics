package main

import (
	"runtime"
	"fmt"
	"time"
    "context"
    "github.com/viciious/go-tarantool"
)

func bTarToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func insertTarantoolData(itemsNum int) {

	var ctx = context.Background()

    opts := tarantool.Options{User: "guest"}
    conn, err := tarantool.Connect("127.0.0.1:3301", &opts)
    if err != nil {
        fmt.Printf("Connection refused: %s\n", err.Error())
	    return
    }

    start := time.Now()

	for i := 0; i < itemsNum; i++ {
		item := newItem(i)

        query := &tarantool.Insert{Space: "examples", Tuple: []interface{}{item.ID, item.Name, item.Description, item.Year}}
        resp := conn.Exec(ctx, query)

        if resp.Error != nil {
            fmt.Println("Insert failed", resp.Error)
        } else {
            fmt.Println(fmt.Sprintf("Insert succeeded: %#v", resp.Data))
        }
		
	}

	conn.Close()
	
	var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v MiB", bRedisToMb(m.Alloc))
    fmt.Printf("\tTotalAlloc = %v MiB", bRedisToMb(m.TotalAlloc))
    fmt.Printf("\tSys = %v MiB", bRedisToMb(m.Sys))
    fmt.Printf("\tNumGC = %v\n", m.NumGC)

	fmt.Println(time.Since(start))

}