package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
	"webApi/internal/config"
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
	cfg := config.GetConfig()

	address := fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Millisecond,
		ReadTimeout:  15 * time.Millisecond,
	}

	log.Infof("server is listening port %s", address)
	log.Fatal(server.Serve(listener))
}
