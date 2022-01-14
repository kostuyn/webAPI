package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
	"webApi/internal/user"
	"webApi/pkg/logging"
)

func main() {
	log := logging.GetLogger()
	log.Info("create router")

	router := httprouter.New()
	userHandler := user.NewHandler(log)
	userHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log := logging.GetLogger()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Millisecond,
		ReadTimeout:  15 * time.Millisecond,
	}

	log.Info("server is listening port 0.0.0.0:1234")
	log.Fatal(server.Serve(listener))
}
