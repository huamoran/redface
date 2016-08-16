package server_test

import (
	"log"

	"github.com/Akagi201/redface/resp"
	"github.com/Akagi201/redface/server"
)

var pongSS = resp.NewRespSimple("PONG")

func pingHandler(args []string) (interface{}, error) {
	return pongSS, nil
}

func Example() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v\n", err)
		}
	}()

	srv, err := server.NewServer("tcp", "127.0.0.1", 6379)
	if err != nil {
		panic(err)
	}

	srv.Handle("ping", pingHandler)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}