package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"imgbot/bot"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	err := bot.Start(token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}