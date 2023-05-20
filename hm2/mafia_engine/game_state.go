package mafia_states

import (
	"fmt"
	"math"
	"sync"
)

type playersState struct {
	players_alive map[string]bool
	players_roles map[string]string
	joined_users  map[string]bool
	connected     map[string]bool
	wait_list     map[string]bool
	cv            sync.Cond
	cv_wait_cnt   int
	cv_cnt        int
	votes         []string
	messages      []string
}

func newPlayersState(users []string) *playersState {
	result := &playersState{players_alive: make(map[string]bool),
		players_roles: make(map[string]string),
		joined_users:  make(map[string]bool),
		votes:         make([]string, 0),
		connected:     make(map[string]bool),
		wait_list:     make(map[string]bool)}
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
	for i, value := range users {
		result.players_alive[value] = true
		result.players_roles[value] = roles[i]
		result.wait_list[value] = true
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
