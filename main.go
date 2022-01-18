package main

import (
	"github.com/hhandhuan/grpc-demo/server"
	"time"
)

func main() {
	go func() {
		server.Run()
	}()
	time.Sleep(time.Hour)
}
