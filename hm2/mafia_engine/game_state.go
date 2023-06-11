package mafia_states

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	graph "hm2/sessions_stat_client"
	"math"
	"sync"
	"time"

	"github.com/Khan/genqlient/graphql"
)

func GetSHA256OfCurrentTime() string {
	currentTime := time.Now().String()
	hasher := sha256.New()
	hasher.Write([]byte(currentTime))
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)

	return hashString
}

type playersState struct {
	players_alive  map[string]bool
	players_roles  map[string]string
	joined_users   map[string]bool
	connected      map[string]bool
	wait_list      map[string]bool
	cv             sync.Cond
	cv_wait_cnt    int
	cv_cnt         int
	votes          []string
	messages       []string
	client_graphql *graphql.Client
	curr_game_id   string
}

func newPlayersState(users []string, client *graphql.Client) *playersState {
	result := &playersState{players_alive: make(map[string]bool),
		players_roles:  make(map[string]string),
		joined_users:   make(map[string]bool),
		votes:          make([]string, 0),
		connected:      make(map[string]bool),
		wait_list:      make(map[string]bool),
		client_graphql: client,
		curr_game_id:   GetSHA256OfCurrentTime()}
	result.cv = *sync.NewCond(&sync.Mutex{})
	result.cv_cnt = len(users)
	result.cv_wait_cnt = 0

	roles := make([]string, 0)
	for i := int(0); i < int(math.Max(1, float64(len(users))*0.25)); i++ {
		roles = append(roles, "mafia")
		roles = append(roles, "commissioner")
	}
	for len(roles) < len(users) {
		roles = append(roles, "man")
	}
	players_gh := make([]graph.PlayerData, 0)
	for i, value := range users {
		players_gh = append(players_gh, graph.PlayerData{Username: value, Role: roles[i]})
		result.players_alive[value] = true
		result.players_roles[value] = roles[i]
		result.wait_list[value] = true
	}

	ctx := context.Background()
	_, err := graph.CreateGame(ctx, *client, graph.GameInfo{Id: result.curr_game_id, State: graph.GameStateInProgress}, players_gh)
	if err != nil {
		fmt.Printf("Graphql error %v\n", err)
	}
	return result
}

func (p *playersState) joinUser(username string) error {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	_, ok := p.joined_users[username]
	if ok {
		return fmt.Errorf("user %v already joined", username)
	}
	p.joined_users[username] = true
	p.connected[username] = true
	return nil
}

func (p *playersState) getRole(username string) (string, error) {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	role, ok := p.players_roles[username]
	if ok {
		return role, nil
	}
	return "", fmt.Errorf("no user %v", username)
}

func (p *playersState) killByVotes() {
	chart := make(map[string]int)
	for _, value := range p.votes {
		alive, ok := p.players_alive[value]
		if ok && alive {
			cnt, ok := chart[value]
			if ok {
				chart[value] = cnt + 1
			} else {
				chart[value] = 1
			}
		}
	}
	p.votes = make([]string, 0)
	if len(chart) == 0 {
		return
	}
	result := ""
	mx := int(0)
	for key, value := range chart {
		if mx < value {
			mx = value
			result = key
		}
	}
	cnt := int(0)
	for _, value := range chart {
		if value == mx {
			cnt += 1
		}
	}
	if cnt == 1 {
		p.players_alive[result] = false
	}
}

func (p *playersState) isAlive(username string) bool {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	result, ok := p.players_alive[username]
	if !ok {
		return false
	}
	return result
}

func (p *playersState) addVote(username string) {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	p.votes = append(p.votes, username)
}

func (p *playersState) getUsersAliveArray() ([]string, []bool) {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	users := make([]string, 0)
	alive := make([]bool, 0)
	for key, value := range p.players_alive {
		users = append(users, key)
		alive = append(alive, value)
	}
	return users, alive
}

func (p *playersState) addMessage(message string) {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	p.messages = append(p.messages, message)
}

func (p *playersState) getMessages() []string {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	result := make([]string, 0)
	result = append(result, p.messages...)
	return result
}

func (p *playersState) isWin() string {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	red := 0
	black := 0
	for user, alive := range p.players_alive {
		if alive {
			if p.players_roles[user] == "mafia" {
				black += 1
			} else {
				red += 1
			}
		}
	}
	if black >= red {
		return "MAFIA WIN."
	}
	if black == 0 {
		return "NOT MAFIA WIN."
	}
	return ""
}

func (p *playersState) waitGamers(username string) error {
	p.cv.L.Lock()
	defer p.cv.L.Unlock()
	_, ok := p.wait_list[username]
	if !ok {
		return fmt.Errorf("no user % v in wait list", username)
	}

	connected := p.connected[username]
	wakeup := (p.cv_wait_cnt+1 == p.cv_cnt)

	if wakeup {
		p.cv.Broadcast()
	} else {
		p.cv_wait_cnt += 1
		if connected {
			p.cv.Wait()
		}
	}

	if wakeup {
		p.killByVotes()
		p.cv_wait_cnt = 0
		p.wait_list = make(map[string]bool)
		for us := range p.connected {
			p.wait_list[us] = true
		}
		p.cv_cnt = len(p.wait_list)

		ctx := context.Background()
		state := graph.GameStateInProgress
		p.cv.L.Unlock()
		if len(p.isWin()) != 0 {
			state = graph.GameStateEnd
		}
		p.cv.L.Lock()
		statuses := make([]graph.PlayerStatus, 0)
		for user, alive := range p.players_alive {
			statuses = append(statuses, graph.PlayerStatus{Username: user, Alive: alive})
		}
		resp, err := graph.UpdateGame(ctx, *p.client_graphql, p.curr_game_id, state, statuses)
		if err != nil || !resp.UpdateGame {
			fmt.Printf("Graphql error %v or ok %v \n", err, resp.UpdateGame)
		}
	}
	return nil
}

func (p *playersState) disconnectUser(username string) {
	p.cv.L.Lock()
	c, ok := p.connected[username]
	if !ok {
		p.cv.L.Unlock()
		return
	}
	if c {
		p.connected[username] = false
		p.players_alive[username] = false
		p.cv.L.Unlock()
		p.waitGamers(username)
		return
	}
	p.cv.L.Unlock()
}
