package main

import (
	"log"

	"github.com/thaddeuscleo/gopetuah/pkg/proxy"
)

func main() {

	server := proxy.New(":3000")
	err := server.Start()
	if err != nil {
		log.Println(err)
	}
}
