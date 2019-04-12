package main

import (
	"github.com/vitorleal/payment/server"
	"runtime"
)

func main() {
	// Max goroutine
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start the API server
	server.Start()
}
