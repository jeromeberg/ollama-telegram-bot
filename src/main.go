package main

import (
	"log"
	"gopkg.in/telebot.v3"
	"time"
	"os"
	"io"
)

type UserContext struct {
	UserID    int64    `json:"user_id"`
	History   []string `json:"history"`
}

func main() {
	// Config
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Logs
	if(config.EnableLog){
		logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error log file: %v", err)
		}
		defer logFile.Close()
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
		log.SetFlags(log.Ldate | log.Ltime)
	}

	// Telegram bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	userContexts := make(map[int64]*UserContext)
	handlers(bot, config, userContexts)

	log.Println("ollama-telegram-bot running...")
	bot.Start()
}
