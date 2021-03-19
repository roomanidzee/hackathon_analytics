package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

        var start = 100000
	var end = 30000000

	
	for i := start; i < end; i+=start {
		var sbRe strings.Builder

		sbRe.WriteString("Loading to Reindexer ")
		sbRe.WriteString(strconv.Itoa(i))
		sbRe.WriteString(" records")

		fmt.Println(sbRe.String())
		insertReindexData(i)
	}

	for i := start; i < end; i+=start {

                var sbR strings.Builder

		sbR.WriteString("Loading to Redis ")
		sbR.WriteString(strconv.Itoa(i))
		sbR.WriteString(" records")

		fmt.Println(sbR.String())
		insertRedisData(i)
	}

	for i := start; i < end; i+=start {

        var sbT strings.Builder

		sbT.WriteString("Loading to Tarantool ")
		sbT.WriteString(strconv.Itoa(i))
		sbT.WriteString(" records")

		fmt.Println(sbT.String())
		insertTarantoolData(i)
	}
}
