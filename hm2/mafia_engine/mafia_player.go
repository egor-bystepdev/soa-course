package mafia_states

import (
	"errors"
	"fmt"
	interact "hm2/interaction"
	game "hm2/pkg/game_proto"
	chat "hm2/queue_chat/chat"
	"math/rand"
	"time"
)

type CitizenGameState struct {
	Username      string
	Role          string
	alive         bool
	day_num       int
	curr_time     string
	players       map[string]bool
	Handgame      bool
	Day_chat      chan string
	Night_chat    chan string
	Stdout_writer *chat.StdoutWriter
}

func (c *CitizenGameState) Start(stream game.Game_JoinClient) error {
	c.day_num = 0
	c.alive = true
	c.players = make(map[string]bool)
	for {
		c.day_num += 1
		c.curr_time = "day"
		c.Stdout_writer.Print(fmt.Sprintf("Process day %v\n", c.day_num))

		exit, err := c.newDayProcess(stream)
		if err != nil {
			fmt.Printf("Day error %v\n", err)
			return err
		}
		if exit {
			return nil
		}
	}
}

func (c *CitizenGameState) newDayProcess(stream game.Game_JoinClient) (bool, error) {
	time.Sleep(time.Second) // For representativeness
	if !c.alive {
		exit, err := c.processStatusOrResult(stream)
		if err != nil {
			return false, err
		}
		if exit {
			return true, nil
		}
		c.curr_time = "night"
		return c.processStatusOrResult(stream)
	}
	if c.day_num != 1 {
		err := c.killSomeone(stream, true)
		if err != nil {
			return false, err
		}
	}
	day_end_msg := &game.GameRequest{GameMessage: &game.GameRequest_EndDay{}}
	err := stream.Send(day_end_msg)
	if err != nil {
		return false, err
	}
	var exit bool
	exit, err = c.processStatusOrResult(stream)
	if err != nil {
		return false, err
	}
	if exit {
		return true, err
	}
	// process night
	c.curr_time = "night"

	time.Sleep(time.Second) // For representativeness
	return c.newNightProcess(stream)
}

func (c *CitizenGameState) killSomeone(stream game.Game_JoinClient, day bool) error {
	users_to_kill := []string{}
	users_to_kill_map := make(map[string]bool)
	for user, alive := range c.players {
		if user == c.Username || !alive {
			continue
		}
		users_to_kill = append(users_to_kill, user)
		users_to_kill_map[user] = true
	}
	if len(users_to_kill) == 0 {
		return errors.New("0 people to kill")
	}
	randomIndex := rand.Intn(len(users_to_kill))
	user_to_kill := users_to_kill[randomIndex]

	if c.Handgame {
		message_handles := make(map[string]func(string))
		if day {
			message_handles["/message"] = c.dayMessage
		} else {
			message_handles["/message"] = c.nightMessage
		}
		user_to_kill = interact.ReadCommand("Choose who you want to kill (you can't not kill, the world is cruel)",
			"/kill",
			users_to_kill_map,
			message_handles,
			c.Stdout_writer)
	} else {
		if day {
			c.dayMessage("bottom text")
		} else {
			c.nightMessage("bottom text")
		}
	}
	kill := &game.GameRequest{GameMessage: &game.GameRequest_KillSomeone{KillSomeone: &game.KillSomeone{Username: user_to_kill}}}
	return stream.Send(kill)
}

func (c *CitizenGameState) investigate(stream game.Game_JoinClient) error {
	users := make(map[string]bool)
	for user, alive := range c.players {
		if user == c.Username || !alive {
			continue
		}
		users[user] = true
	}
	if len(users) == 0 {
		return errors.New("0 people to kill")
	}
	us := ""
	for u := range users {
		us = u
		break // Прерываем цикл после первой итерации, чтобы получить "первый" элемент
	}
	if c.Handgame {
		us = interact.ReadCommand("Choose who you want to check", "/check", users, map[string]func(string){}, c.Stdout_writer)
	}
	inv := &game.GameRequest{GameMessage: &game.GameRequest_Investigate{Investigate: &game.Investigate{Username: us}}}
	if err := stream.Send(inv); err != nil {
		return err
	}

	msg, err := stream.Recv()
	if err != nil {
		return err
	}

	switch msg.GameMessage.(type) {
	case *game.GameResponse_InvestigateResult:
		is := msg.GetInvestigateResult().GetMafia()
		if is {
			c.Stdout_writer.Print(fmt.Sprintf("%v is mafia\n", us))
		} else {
			c.Stdout_writer.Print(fmt.Sprintf("%v is not mafia\n", us))
		}
		if is {
			result := ""
			if c.Handgame {
				result = interact.ReadCommand("publish this information?",
					"/publish",
					map[string]bool{"yes": true, "no": true},
					map[string]func(string){}, c.Stdout_writer)
			}
			pub := &game.GameRequest{GameMessage: &game.GameRequest_InvestigatePublish{InvestigatePublish: &game.InvestigatePublish{Answer: (result == "yes")}}}
			if err := stream.Send(pub); err != nil {
				return err
			}
		}
	default:
		return errors.New("invalid msg for com")
	}
	return nil
}

func (c *CitizenGameState) newNightProcess(stream game.Game_JoinClient) (bool, error) {
	if !c.alive {
		return c.processStatusOrResult(stream) // TODO error hand
	}
	if c.Role == "mafia" && c.day_num != 1 {
		err := c.killSomeone(stream, false)
		if err != nil {
			return false, err
		}
	}
	if c.Role == "commissioner" && c.day_num != 1 {
		err := c.investigate(stream)
		if err != nil {
			return false, err
		}
	}
	day_end_msg := &game.GameRequest{GameMessage: &game.GameRequest_EndDay{}}
	err := stream.Send(day_end_msg)
	if err != nil {
		return false, err
	}

	exit, err := c.processStatusOrResult(stream)

	if err != nil {
		return false, err
	}
	if exit {
		return true, err
	}

	return false, nil
}

func (c *CitizenGameState) processStatusOrResult(stream game.Game_JoinClient) (bool, error) {
	msg, err := stream.Recv()
	if err != nil {
		return false, err
	}
	switch msg.GameMessage.(type) {
	case *game.GameResponse_GameStatus:
		err := c.handleUpdateStatus(msg.GetGameStatus())
		if err != nil {
			return false, err
		}
	case *game.GameResponse_Result:
		return c.handleResult(msg.GetResult())
	default:
		fmt.Printf("Unexpected message, type %T. Exit.\n", msg.GameMessage)
		return false, errors.New("unknown message")
	}
	return false, nil
}

func (c *CitizenGameState) handleResult(result *game.Result) (bool, error) {
	c.Stdout_writer.Print(fmt.Sprintf("Game result %v\n!!!", result.Result))
	return true, nil
}

func (c *CitizenGameState) handleUpdateStatus(status *game.GameStatus) error {
	usernames := status.Usernames
	alive := status.Alive
	day_num := status.DayNum
	time := status.Time
	messages := status.Messages
	if len(usernames) != len(alive) {
		s := fmt.Sprintf("`alive` size %v not equal to `usernames` size %v", len(alive), len(usernames))
		return errors.New(s)
	}
	if day_num != int32(c.day_num) {
		s := fmt.Sprintf("respose day num %v not equal to client day num %v", day_num, c.day_num)
		return errors.New(s)
	}
	if time != c.curr_time {
		s := fmt.Sprintf("respose time  %v not equal to client time %v", time, c.curr_time)
		return errors.New(s)
	}
	c.Stdout_writer.Print(fmt.Sprintf("Day num %v, time %v\n", c.day_num, c.curr_time))

	state_str := ""
	for i := range usernames {
		c.players[usernames[i]] = alive[i]
		substr := ""
		if usernames[i] == c.Username {
			substr = "(YOU) "
		}
		alive_status := "(DEAD) "
		if alive[i] {
			alive_status = "IN GAME"
		}
		state_str += fmt.Sprintf("%v %v-- %v\n", usernames[i], substr, alive_status)
	}

	c.Stdout_writer.Print(state_str)
	if len(messages) != 0 {
		msgs_str := "Messages\n"
		c.Stdout_writer.Print("Messages\n")
		for _, m := range messages {
			msgs_str += m + "\n"
		}
		c.Stdout_writer.Print(msgs_str)
	}
	c.alive = c.players[c.Username]
	return nil
}

func (c *CitizenGameState) dayMessage(message string) {
	select {
	case c.Day_chat <- message:
	default:
	}
}

func (c *CitizenGameState) nightMessage(message string) {
	select {
	case c.Night_chat <- message:
	default:
	}
}
