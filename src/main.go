package main

import (
	"context"
	"kmsbot/bootstrap"
	"kmsbot/domain"
	"kmsbot/rest"
	"kmsbot/service"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	config, err := bootstrap.New()
	if err != nil {
		log.Fatalln("fatal", err)
	}

	botService, err := service.NewBot(config.Bot.Token)
	if err != nil {
		log.Fatalln("bot service", err)
	}

	mikrotikService, err := service.NewMikrotik(config.Router.Addr, config.Router.User, config.Router.Password)
	if err != nil {
		log.Fatalln("mikrotik service", err)
	}

	storeService, err := service.NewStore(config.DB.Name, config.DB.User, config.DB.Password)
	if err != nil {
		log.Fatalln("store service", err)
	}

	core := domain.NewCore(botService, storeService, mikrotikService)
	go core.Start()

	srv := rest.NewServer(config.Server.Port, core)
	srv.SetRoutes()

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != net.ErrClosed {
			log.Fatalln("listen and serve", err)
		}
	}()

	<-shutdown

	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Println("srv shutdown", err)
	}
}
