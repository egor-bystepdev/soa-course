package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
	"unicode"

	game "hm2/mafia_engine"
	welcome "hm2/pkg/welcome_proto"

	"google.golang.org/grpc"
)

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type server struct {
	welcome.UnimplementedWelcomeServer

	users              map[string]int
	sessions           map[string]int
	session_to_port    map[int]int
	started_sessions   map[int]bool
	current_session_id int
	session_ports      int
	mutex              sync.Mutex
}

func deleteUser(s *server, username string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.users, username)
}

func getUsersBySession(s *server, session_id int) []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	result := make([]string, 0)
	for key, value := range s.sessions {
		if value == session_id {
			result = append(result, key)
		}
	}
	return result
}

func checkAndSetSession(s *server, username string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.users[username]; !ok {
		return s.sessions[username]
	} else if len(s.users) >= kGameUsersLimit {
		s.current_session_id += 1
		s.session_ports += 1
		s.session_to_port[s.current_session_id] = s.session_ports
		i := 0
		var keys []string
		for key := range s.users {
			if i >= kGameUsersLimit {
				break
			}
			s.sessions[key] = s.current_session_id
			keys = append(keys, key)
			i += 1
		}
		for _, value := range keys {
			delete(s.users, value)
		}
		return s.current_session_id
	}

	return 0
}

func selectSessionData(s *server, session_id int) (string, int, int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	m := map[string]interface{}{
		"port": s.session_to_port[session_id],
		"id":   session_id,
	}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Json serialize error %v\n", err)
		return "", 0, 0
	}
	jsonString := string(jsonBytes)
	return jsonString, s.session_to_port[session_id], session_id
}

func canStartSession(s *server, session_id int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	started, ok := s.started_sessions[session_id]
	s.started_sessions[session_id] = true
	if !ok {
		return true
	}
	return !started
}

func sendList(s *server, stream welcome.Welcome_ConnectServer) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var keys []string
	for k := range s.users {
		keys = append(keys, k)
	}
	jsonBytes, err := json.Marshal(keys)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return false
	}
	jsonString := string(jsonBytes)
	err = stream.Send(&welcome.WelcomeResponse{MessageType: welcome.MessageType_WaitList, ResultString: jsonString})
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return false
	}
	return true
}

var kGameUsersLimit int = 4

func (s *server) Connect(req *welcome.ConnectRequest, stream welcome.Welcome_ConnectServer) error {
	username := req.GetUsername()
	fmt.Printf("Received message: %s\n", username)

	var result_type welcome.MessageType
	var result_string = ""
	if isAlphanumeric((username)) {
		s.mutex.Lock()
		var ok bool
		_, ok = s.users[username]
		if !ok {
			s.users[username] = 0
		} else {
			result_type = welcome.MessageType_RepeatedName
			result_string = "this username in use."
		}
		s.mutex.Unlock()
	} else {
		result_type = welcome.MessageType_InvalidName
		result_string = "username must contain only letters and digits."
	}
	if len(result_string) != 0 {
		stream.Send(&welcome.WelcomeResponse{MessageType: result_type, ResultString: result_string})
		return nil
	}
	if !sendList(s, stream) {
		deleteUser(s, username)
		return nil
	}
	for {
		maybe_session := checkAndSetSession(s, username)
		if maybe_session != 0 {
			// go game.StartGameServer(maybe_session, )
			json_data, port, id := selectSessionData(s, maybe_session)

			log.Printf("Start session1111")
			if canStartSession(s, id) {
				log.Printf("Start session")
				go game.StartGameServer(id, port, getUsersBySession(s, id))
			}
			time.Sleep(time.Second)
			stream.Send(&welcome.WelcomeResponse{MessageType: welcome.MessageType_SessionConnectAddr, ResultString: json_data})
			return nil
		}
		if !sendList(s, stream) {
			deleteUser(s, username)
			return nil
		}
		time.Sleep(time.Second)
	}
}

func main() {
	envPort := os.Getenv("MAFIA_SERVER_PORT")
	envGamePlayers := os.Getenv("MAFIA_PLAYERS_COUNT")
	if len(envGamePlayers) != 0 {
		game_players, err := strconv.Atoi(envGamePlayers)
		if err == nil && game_players >= 3 {
			kGameUsersLimit = game_players
		} else {
			fmt.Printf("Invalid MAFIA_PLAYERS_COUNT %v", envGamePlayers)
			return
		}
	}
	if len(envPort) == 0 {
		envPort = "50051"
	}

	fmt.Printf("Start server with %v players, port %v\n", kGameUsersLimit, envPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", envPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	welcome.RegisterWelcomeServer(s, &server{users: make(map[string]int),
			sessions: make(map[string]int),
			session_to_port: make(map[int]int),
			started_sessions: make(map[int]bool), 
			session_ports: 5000,
			current_session_id: 0})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
