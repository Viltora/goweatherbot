package handlers

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	var sb strings.Builder

	sb.WriteString("<b>Расшифровка таблицы:</b>\n")
	sb.WriteString("• <b>TEMP</b> — Температура воздуха (°C)\n")
	sb.WriteString("• <b>PREC</b> — Осадки за час (мм водяного столба)\n")
	sb.WriteString("• <b>WIND</b> — Скорость ветра (м/с) и направление\n")
	sb.WriteString("• <b>SNOW</b> — Общая высота снега на земле (см)\n")
	sb.WriteString("• <b>VPD</b> — Дефицит давления пара (кПа). Чем выше показатель, тем суше .\n\n")

	sb.WriteString("<b>Особенности:</b>\n")
	sb.WriteString("Данные обновляются в реальном времени.\n\n")
	sb.WriteString("Прогноз разбит на блоки по 6 часов для удобства (ночь, утро, день, вечер).\n\n")

	sb.WriteString("<i>Бот создан в некоммерческих целях с использованием открытых данных Open-Meteo.</i>\n\n")
	sb.WriteString("<i>Для обратной связи: @Viltora.</i>\n\n")

	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      sb.String(),
		ParseMode: models.ParseModeHTML,
	})
}
