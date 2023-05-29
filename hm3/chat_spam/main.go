package main

import (
	"fmt"
	chat "hm3/chat"
	"time"

	"math/rand"
	"sync"
)

func RandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func main() {
	url := "amqp://guest:guest@localhost:5672/"
	chat.CreateFanoutExchange(url, "test")
	writer := chat.StdoutWriter{Mutex: sync.Mutex{}}
	username := RandomString(10)
	fmt.Printf("You %v\n", username)
	messages := make(chan string)
	go chat.Publish(url, "test", username, messages)
	go chat.Consume(url, "test", username, &writer)

	i := 0
	for {
		messages <- fmt.Sprintf("Its my message number %v, my random text %v", i, RandomString(10))
		i++
		time.Sleep(time.Second)
	}
}
