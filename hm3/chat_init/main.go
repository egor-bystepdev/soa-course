package main

import (
	"fmt"
	chat "hm3/chat"

	"log"
)

func main() {
	url := "amqp://guest:guest@localhost:5672/"
	err := chat.CreateFanoutExchange(url, "test")

	if err != nil {
		fmt.Printf("Err %v", err)
		return
	}

	log.Printf("To exit press CTRL+C")
	for {

	}
}
