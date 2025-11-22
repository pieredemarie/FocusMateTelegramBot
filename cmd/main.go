package main

import (
	"focusMate/internal/bot"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    token := os.Getenv("TELEGRAM_TOKEN")
    if token == "" {
        log.Fatal("TELEGRAM_TOKEN is not set in .env")
    }

    pref := telebot.Settings{
        Token:  token,
        Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
    }

    b, err := telebot.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }

    b.Handle("/remind", func(c telebot.Context) error {
		bot.HandleRemind(c, b) 
		return nil
	})

    log.Println("Bot started...")
    b.Start()
}