package game_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	game "hm2/pkg/game_proto"

	gamer "hm2/mafia_engine"
)

func Start(session_id, port int, username string, manual_game bool) error {

	log.Println("Hello!!!!")
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
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
	g := gamer.CitizenGameState{Username: username, Role: role, Handgame: manual_game}
	err = g.Start(stream)
	if err != err {
		fmt.Printf("Error occured %v.", err)
	}
	return nil
}
