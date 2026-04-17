package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"goweatherbot/handlers"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN не найден")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	handlers.SetBotCommands(b)

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/start",
		bot.MatchTypeExact,
		handlers.StartHandler,
	)

	// Список команд для локаций
	openMeteoCommands := map[string]string{
		"/rosapeak_openmeteo":       "rosapeak",
		"/rosa1600_openmeteo":       "rosa1600",
		"/caucaseexpress_openmeteo": "caucaseexpress",
		"/krokus_openmeteo":         "krokus",
		"/edelweiss_openmeteo":      "edelweiss",
	}

	for command, locationCode := range openMeteoCommands {
		loc := handlers.Locations[locationCode]

		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			command,
			bot.MatchTypeExact,
			handlers.OpenMeteoHandler(
				loc.Latitude,
				loc.Longitude,
				loc.Elevation,
				loc.Name,
			),
		)
	}

	log.Println("🤖 Weather bot запущен (режим: почасовой на сегодня)")
	b.Start(ctx)
}
