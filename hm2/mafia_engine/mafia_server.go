package mafia_states

import (
	"fmt"
	game "hm2/pkg/game_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	game.UnimplementedGameServer
	state      playersState
	session_id int
}

func sendInfo(s *server, stream game.Game_JoinServer, day_num int, time string) error {
	users, alive := s.state.getUsersAliveArray()
	messages := s.state.getMessages()
	if len(messages) != 0 {
		fmt.Printf("M %v\n", messages)
	}
	info := &game.GameResponse{GameMessage: &game.GameResponse_GameStatus{GameStatus: &game.GameStatus{Usernames: users, Alive: alive,
		DayNum: int32(day_num), Time: time, Messages: messages}}}
	return stream.Send(info)
}

func handleInvestigationPlayer(s *server, stream game.Game_JoinServer, username string) error {
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	inv_username := ""
	switch msg.GameMessage.(type) {
	case *game.GameRequest_Investigate:
		inv_username = msg.GetInvestigate().GetUsername()
	default:
		return fmt.Errorf("not expected connect message")
	}
	role, err := s.state.getRole(inv_username)
	if err != nil {
		return err
	}
	result_ans := (role == "mafia")
	result_message := &game.GameResponse{GameMessage: &game.GameResponse_InvestigateResult{InvestigateResult: &game.InvestigateResult{Mafia: result_ans}}}
	stream.Send(result_message)
	if result_ans {
		msg, err = stream.Recv()
		if err != nil {
			return err
		}
		publish := false
		switch msg.GameMessage.(type) {
		case *game.GameRequest_InvestigatePublish:
			publish = msg.GetInvestigatePublish().GetAnswer()
			if publish {
				s.state.addMessage(fmt.Sprintf("%v says that %v mafia", username, inv_username))
			}
		default:
			return fmt.Errorf("not expected connect message")
		}
	}

	{
		fmt.Println("Wait.11111")
		msg, err := stream.Recv()
		fmt.Println("Wait.22222")
		if err != nil {
			return err
		}
		switch msg.GameMessage.(type) {
		case *game.GameRequest_EndDay:
		default:
			return fmt.Errorf("expected end day message")
		}
	}
	fmt.Printf("Wait gamers %v\n", username)
	s.state.waitGamers(username)
	return nil
}

func handleEndDayOrKill(s *server, stream game.Game_JoinServer, username string, expected_kill bool) error {
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	is_day_end := false
	kill_username := ""
	switch msg.GameMessage.(type) {
	case *game.GameRequest_EndDay:
		is_day_end = true
	case *game.GameRequest_KillSomeone:
		kill_username = msg.GetKillSomeone().GetUsername()
	default:
		return fmt.Errorf("not expected connect message")
	}
	if !is_day_end && expected_kill {
		s.state.addVote(kill_username)
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		switch msg.GameMessage.(type) {
		case *game.GameRequest_EndDay:
			is_day_end = true
		default:
			return fmt.Errorf("expected end day message")
		}
	}
	s.state.waitGamers(username)
	return nil
}

func winExit(s *server, stream game.Game_JoinServer) bool {
	result := s.state.isWin()
	if len(result) == 0 {
		return false
	}
	result_message := &game.GameResponse{GameMessage: &game.GameResponse_Result{Result: &game.Result{Result: result}}}
	stream.Send(result_message)
	return true
}

func processGame(s *server, stream game.Game_JoinServer, username string) error {
	day_num := int(1)
	role, err := s.state.getRole(username)
	if err != nil {
		return err
	}
	for {
		fmt.Printf("Session id %v, day %v\n", s.session_id, day_num)

		{ // DAY
			if s.state.isAlive(username) {
				if err := handleEndDayOrKill(s, stream, username, (day_num != 1)); err != nil {
					return err
				}
			} else {
				s.state.waitGamers(username)
			}

			if winExit(s, stream) {
				return nil
			}

			fmt.Printf("Session id %v, send day info\n", s.session_id)
			if err := sendInfo(s, stream, day_num, "day"); err != nil {
				return err
			}
		}

		{ // NIGHT
			fmt.Printf("Session id %v, wait night\n", s.session_id)

			if s.state.isAlive(username) {
				if role == "commissioner" && day_num != 1 {
					err := handleInvestigationPlayer(s, stream, username)
					if err != nil {
						return err
					}
				} else if err := handleEndDayOrKill(s, stream, username, (role == "mafia" && day_num != 1)); err != nil {
					return err
				}
			} else {
				s.state.waitGamers(username)
			}

			if winExit(s, stream) {
				return nil
			}

			fmt.Printf("Session id %v, send night info\n", s.session_id)
			if err := sendInfo(s, stream, day_num, "night"); err != nil {
				return err
			}
		}

		day_num += 1
	}
}

func (s *server) Join(stream game.Game_JoinServer) error {
	log.Println("Started join.")
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	username := ""
	session_id := 0
	switch msg.GameMessage.(type) {
	case *game.GameRequest_GameConnect:
		username = msg.GetGameConnect().GetUsername()
		session_id = int(msg.GetGameConnect().GetSessionId())
	default:
		return fmt.Errorf("expected connect message")
	}
	if session_id != s.session_id {
		return fmt.Errorf("invalid session id %v", session_id)
	}
	s.state.joinUser(username)
	fmt.Println("Wait.")
	if err := s.state.waitGamers(username); err != nil {
		return err
	}

	role, err := s.state.getRole(username)
	if err != nil {
		return err
	}

	log.Println("Send role.")
	role_msg := &game.GameResponse{GameMessage: &game.GameResponse_GameRole{GameRole: &game.GameRole{Role: role}}}
	if err := stream.Send(role_msg); err != nil {
		return err
	}

	err = processGame(s, stream, username)
	if err != nil {
		fmt.Printf("Error occured %v\n", err)
		s.state.disconnectUser(username)
	}
	return err
}

func StartGameServer(session_id, port int, users []string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	log.Printf("Start server with session %v, port %v", session_id, port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	game.RegisterGameServer(s, &server{session_id: session_id, state: *newPlayersState(users)})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
