package graph

import (
	"hm2/sessions_stat_service/graph/model"
	"sync"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	Games []model.Game
	Index map[string]int
	Mutex sync.Mutex
}
