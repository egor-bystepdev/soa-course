package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
	"math/rand"
	"time"
	"bufio"
	"encoding/json"

	"google.golang.org/grpc"

	game_client "hm2/game"
	"hm2/pkg/welcome_proto"
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

func ParseWaitList(json_value string) ([]string, error) {
	var users []string
    err := json.Unmarshal([]byte(json_value), &users)
	return users, err
}

type Session struct {
    Port  int `json:"port"`
    Id   int   `json:"id"`
}

func ParseNewSession(json_value string) (int, int, error) {
	var session Session
    err := json.Unmarshal([]byte(json_value), &session)
	return session.Id, session.Port, err
}

func connectAndPlay(host_port_from_stdin, manual_game bool) error {
	envPort := os.Getenv("MAFIA_SERVER_PORT")
	envHost := os.Getenv("MAFIA_SERVER_HOST")
	if len(envPort) == 0 {
		envPort = "50051"
	}

	if host_port_from_stdin {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Введите хост: ")
		scanner.Scan()
		envHost = scanner.Text()

		fmt.Print("Введите порт: ")
		scanner.Scan()
		envPort = scanner.Text()
	}
	rabbitmq_url := "amqp://guest:guest@localhost:5672/"
	if len(os.Getenv("RABBITMQ_HOST")) != 0 {
		rabbitmq_url = fmt.Sprintf("amqp://guest:guest@%v:5672/", os.Getenv("RABBITMQ_HOST"))
	}
	fmt.Printf("Connect to %v:%v\n", envHost, envPort)
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", envHost, envPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := welcome.NewWelcomeClient(conn)
	username := RandomString(15)
	stream, err := c.Connect(context.Background(), &welcome.ConnectRequest{Username: username})
	if err != nil {
		log.Fatalf("Error calling Connect: %v", err)
		return err
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Connect(_) error %v", err)
		}
		fmt.Printf("Receive %v\n", string(response.ResultString))
		switch response.MessageType {
		case welcome.MessageType_InvalidName:
			fmt.Printf("Invalid name %v\n", response.ResultString)
		case welcome.MessageType_RepeatedName:
			fmt.Printf("Repeated name %v\n", response.ResultString)
		case welcome.MessageType_WaitList:
			users, err := ParseWaitList(response.ResultString)
			if err != nil {
				fmt.Printf("Failed to parse wait list `%v`\n. exit", err)
				return err
			}
			fmt.Println("Wait list:")
			for _, value := range users {
				fmt.Println(value)
			}
			fmt.Println("-----------")
		case welcome.MessageType_SessionConnectAddr:
			id, port, err := ParseNewSession(response.ResultString)
			if err != nil {
				fmt.Printf("Failed to parse %v\n", err)
				return err
			}
			fmt.Printf("Port `%v`, session_id `%v`\n", port, id)
			stream.CloseSend()
			err = game_client.Start(id, envHost, port, username, manual_game, rabbitmq_url)
			if err != nil {
				return err
			}
			fmt.Println("Game end. Thanks.")
			time.Sleep(time.Second)
		}
	}
	return nil
}

func main() {
	// Определение параметра запуска
	pFlag := flag.Bool("p", false, "Manual entry of the address")
	hFlag := flag.Bool("h", false, "Manual game")

	// Парсинг параметров запуска
	flag.Parse()
	for {
		err := connectAndPlay(*pFlag, *hFlag)
		if err != nil {
			fmt.Printf("Error occured %v\n", err)
		}
	}
}
