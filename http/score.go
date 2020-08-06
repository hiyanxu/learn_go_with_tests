package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores   map[string]string
	winCalls []string
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, score)
}

//func PlayerServer(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.URL.Path)
//	player := r.URL.Path[len("/players/"):]
//
//	score := GetPlayerScore(player)
//	fmt.Fprintf(w, score)
//}

func (s PlayerServer) GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
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
	store map[string]string
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) string {
	fmt.Println("hehe")
	return "123"
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	// 将PlayerServer转换为HandlerFunc
	//handler := http.HandlerFunc(PlayerServer)
	handler := &PlayerServer{store: &InMemoryPlayerStore{}}

	// ListenAndServe()监听一个端口，在该端口上处理请求
	if err := http.ListenAndServe(":9091", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
