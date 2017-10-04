package main

import (
	"github.com/ingresse/payment/server"
	"runtime"
)

func main() {
	// Max goroutine
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start the API server
	server.Start()
}
