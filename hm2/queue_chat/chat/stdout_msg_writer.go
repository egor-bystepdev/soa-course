package chat

import (
	"fmt"

	"sync"
)

type StdoutWriter struct {
	Mutex sync.Mutex 
}

func (s *StdoutWriter) Print(message string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	fmt.Println(message)
}

func (s *StdoutWriter) PrintUserMessage(message string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	cian := "\033[36m"
    reset := "\033[0m"
	fmt.Printf("%v%v%v\n", cian, message, reset)
}
