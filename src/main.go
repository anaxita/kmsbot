package main

import (
	"context"
	"io"
	"kmsbot/bootstrap"
	"kmsbot/domain"
	"kmsbot/rest"
	"kmsbot/service"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	f, err := os.Create("./logs/" + time.Now().Format("02-01-2006-15-04-05"+".log"))
	if err != nil {
		log.Fatalln("create logfile", err)
	}

	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))

	shutdown := make(chan os.Signal, 1)
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

	mikrotikService2, err := service.NewMikrotik(config.Router2.Addr, config.Router2.User, config.Router2.Password)
	if err != nil {
		log.Fatalln("mikrotik service", err)
	}

	storeService, err := service.NewStore(config.DB.Name, config.DB.User, config.DB.Password)
	if err != nil {
		log.Fatalln("store service", err)
	}

	core := domain.NewCore(botService, storeService, mikrotikService, mikrotikService2)
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
