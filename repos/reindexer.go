package repos

import (
	"runtime"
	"fmt"
	"time"
	"github.com/restream/reindexer"
	_ "github.com/restream/reindexer/bindings/builtin"
)

func bReToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func insertReindexData(itemsNum int) {

	db := reindexer.NewReindex("cproto://127.0.0.1:6534/testdb", reindexer.WithCreateDBIfMissing())

	db.OpenNamespace("items", reindexer.DefaultNamespaceOptions(), Item{})

	start := time.Now()

	for i := 0; i < itemsNum; i++ {
		err := db.Upsert("items", newItem(i))
		if err != nil {
			panic(err)
		}
	}

	var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v MiB", bReToMb(m.Alloc))
    fmt.Printf("\tTotalAlloc = %v MiB", bReToMb(m.TotalAlloc))
    fmt.Printf("\tSys = %v MiB", bReToMb(m.Sys))
    fmt.Printf("\tNumGC = %v\n", m.NumGC)

	fmt.Println(time.Since(start))

}
