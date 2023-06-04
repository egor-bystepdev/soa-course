package game_client

import (
	"context"
	"fmt"
	"log"
	"time"
	"sync"

	"google.golang.org/grpc"

	game "hm2/pkg/game_proto"

	gamer "hm2/mafia_engine"

	chat "hm2/queue_chat/chat"
)

func Start(session_id int, host string, port int, username string, manual_game bool, rabbitmq_url string) error {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	log.Println("Create client.")

	c := game.NewGameClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	stream, err := c.Join(ctx)
	if err != nil {
		log.Fatalf("Error calling Connect: %v", err)
		return err
	}
	join_msg := &game.GameRequest{GameMessage: &game.GameRequest_GameConnect{GameConnect: &game.GameConnect{SessionId: int32(session_id), Username: username}}}
	stream.Send(join_msg)

	log.Println("Send join.")
	msg, err := stream.Recv()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	var role string
	switch msg.GameMessage.(type) {
	case *game.GameResponse_GameRole:
		role = msg.GetGameRole().Role
	default:
		fmt.Println("Expected role message. Exit.")
		return err
	}

	log.Printf("Role %v.\n", role)
	stdout_writer := chat.StdoutWriter{Mutex: sync.Mutex{}}
	go chat.Consume(rabbitmq_url, fmt.Sprintf("day_%v", session_id), username, &stdout_writer)
	day_ch := make(chan string)
	night_ch := make(chan string)
	go chat.Publish(rabbitmq_url, fmt.Sprintf("day_%v", session_id), username, day_ch)
	if (role == "mafia") {
		go chat.Consume(rabbitmq_url, fmt.Sprintf("night_%v", session_id), username, &stdout_writer)
		go chat.Publish(rabbitmq_url, fmt.Sprintf("night_%v", session_id), username, night_ch)
	}
	g := gamer.CitizenGameState{Username: username, Role: role, Handgame: manual_game, Day_chat: day_ch, Night_chat: night_ch, Stdout_writer: &stdout_writer}
	time.Sleep(time.Second)
	err = g.Start(stream)
	chat.DeleteQueue(rabbitmq_url, fmt.Sprintf("night_%v", session_id))
	chat.DeleteQueue(rabbitmq_url, fmt.Sprintf("day_%v", session_id))
	if err != nil {
		fmt.Printf("Error occured %v.", err)
	}
	return nil
}
