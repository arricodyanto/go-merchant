package main

import (
	"go-merchant/delivery"
)

func main() {
	// run server dari file server.go
	delivery.NewServer().Run()
}
