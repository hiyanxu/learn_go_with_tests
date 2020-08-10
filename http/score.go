package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	//router *http.ServeMux // 该server具有n个路由，可以提前装载好
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store,
		http.NewServeMux(),
	}

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router
	return p
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

type Player struct {
	Name string
	Wins int
}

// 不再需要该方法，因为在http包中，mux多路复用器已经实现了http.Handler接口
//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	//router := http.NewServeMux()
//	////router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	////	//w.WriteHeader(http.StatusOK)
//	////}))
//	//router.Handle("/league", http.HandlerFunc(p.leagueHandler))
//	//
//	//router.Handle("/players/", http.HandlerFunc(p.playerHandler))
//
//	// 路由已经装载好，请求过来直接分发
//	p.router.ServeHTTP(w, r)
//}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")

	//leagues := p.getLeagues()
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

// 当前的请求需要获取相应的json
func (p *PlayerServer) getLeagues() []Player {
	return []Player{
		{
			"Chris",
			20,
		},
	}
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	//player := r.URL.Path[len("/players/"):]
	//score := p.store.GetPlayerScore(player)
	//if score == "" {
	//	w.WriteHeader(http.StatusNotFound)
	//}
	//fmt.Fprintf(w, score)
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	//player := r.URL.Path[len("/players/"):]
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	//player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

//func PlayerServer(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.URL.Path)
//	player := r.URL.Path[len("/players/"):]
//
//	score := GetPlayerScore(player)
//	fmt.Fprintf(w, score)
//}

func (s PlayerServer) GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}

	return 0
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

/**
http handler实现：
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
任何实现了该接口的结构体都可以作为http handler

HandlerFunc是一个函数类型，该类型同样实现了上面的handler接口
*/

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}

func main() {
	// 将PlayerServer转换为HandlerFunc
	//handler := http.HandlerFunc(PlayerServer)
	handler := &PlayerServer{store: NewInMemoryPlayerStore()}

	// ListenAndServe()监听一个端口，在该端口上处理请求
	if err := http.ListenAndServe(":9091", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
