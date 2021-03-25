package main

import (
	"log"
	"net/http"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`Go语言编程之旅`))
	})

	_ = http.ListenAndServe(":6060", http.DefaultServeMux)
}
