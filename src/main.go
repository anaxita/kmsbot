package main

import (
	"kmsbot/bootstrap"
	"kmsbot/domain"
	"kmsbot/service"
	"log"
)

type user struct {
	Name int `json:"name"`
}

func main() {
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

	core.Start()
}
