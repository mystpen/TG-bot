package main

import (
	"log"

	"github.com/mystpen/TG-bot/internal/clients/tg"
	"github.com/mystpen/TG-bot/internal/config"
	"github.com/mystpen/TG-bot/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil{
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}