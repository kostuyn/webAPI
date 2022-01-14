package main

import (
	"log"
	"net"
	"net/http"
	"time"
	"webApi/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	userHandler := user.NewHandler()
	userHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Millisecond,
		ReadTimeout:  15 * time.Millisecond,
	}

	log.Println("server is listening port 0.0.0.0:1234")
	log.Fatalln(server.Serve(listener))
}
