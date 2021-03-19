package main

import (
	"github.com/roomanidze/hackathon_analytics/repos"
)

func main() {
	insertRedisData(100)
	insertTarantoolData(100)
	insertReindexData(100)
}