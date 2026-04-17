package handlers

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SetBotCommands(b *bot.Bot) {
	params := &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{Command: "start", Description: "Приветствие бота и описание команд"},

			{Command: "rosapeak_openmeteo", Description: "Роза Пик 2320м · Open-Meteo"},
			{Command: "rosa1600_openmeteo", Description: "Роза 1600м · Open-Meteo"},
			{Command: "caucaseexpress_openmeteo", Description: "Кавказский экспресс 1350м · Open-Meteo"},
			{Command: "krokus_openmeteo", Description: "Крокус 2509м · Open-Meteo"},
			{Command: "edelweiss_openmeteo", Description: "Эдельвейс 1472м · Open-Meteo"},
		},
	}

	_, err := b.SetMyCommands(context.Background(), params)
	if err != nil {
		log.Println("Ошибка установки команд:", err)
	} else {
		log.Println("Команды успешно установлены ✅")
	}
}
