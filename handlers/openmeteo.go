package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const OpenMeteoBaseURL = "https://api.open-meteo.com/v1/forecast"

var defaultHourlyParams = []string{
	"temperature_2m",
	"precipitation",
	"wind_speed_10m",
	"wind_direction_10m",
	"uv_index",
	"snow_depth",
	"freezing_level_height",
	"lifted_index",
	"vapour_pressure_deficit",
}

func OpenMeteoHandler(lat, lon float64, elevation int, locationName string) func(context.Context, *bot.Bot, *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		apiUrl := buildOpenMeteoURL(lat, lon, elevation)
		resp, err := http.Get(apiUrl)
		if err != nil {
			log.Printf("[ОШИБКА API] %v", err)
			return
		}
		defer resp.Body.Close()

		var data OpenMeteoResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			log.Printf("[ОШИБКА JSON] %v", err)
			return
		}

		text := FormatWeatherResponse(data, locationName)

		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      text,
			ParseMode: models.ParseModeHTML,
		})

		if err != nil {
			log.Printf("[ОШИБКА TG] %v", err)
			// Если сообщение слишком длинное, можно попробовать отправить без HTML
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Ошибка отправки данных. Возможно, слишком много данных.",
			})
		}
	}
}

func FormatWeatherResponse(data OpenMeteoResponse, locationName string) string {
	if len(data.Hourly.Time) == 0 {
		return "Нет данных для отображения"
	}

	getWindDir := func(deg int) string {
		index := int((float64(deg)+22.5)/45.0) % 8
		directions := []string{"С ", "СВ", "В ", "ЮВ", "Ю ", "ЮЗ", "З ", "СЗ"}
		return directions[index]
	}

	dateStr := data.Hourly.Time[0][:10]
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<b>Локация: %s</b>\n", locationName))
	sb.WriteString(fmt.Sprintf("<b>Дата: %s</b>\n\n", dateStr))

	sb.WriteString("<pre>")
	// Добавили колонку VPD
	// TIME  TEMP  PREC WIND DIR SNOW  VPD
	sb.WriteString("TIME  TEMP  PREC WIND DIR SNOW  VPD\n")
	sb.WriteString("-----------------------------------\n")

	for i := 0; i < len(data.Hourly.Time); i++ {
		timeParts := strings.Split(data.Hourly.Time[i], "T")
		timeOnly := "00:00"
		if len(timeParts) > 1 {
			timeOnly = timeParts[1]
		}

		windMS := data.Hourly.WindSpeed[i] / 3.6
		windDirStr := getWindDir(data.Hourly.WindDir[i])

		// Форматируем строку:
		// %4.2f для VPD, так как это значения обычно от 0.00 до 2.00
		row := fmt.Sprintf("%-5s %5.1f %5.1f %4.1f %-2s %4.0f %5.2f\n",
			timeOnly,
			data.Hourly.Temperature[i],
			data.Hourly.Precipitation[i],
			windMS,
			windDirStr,
			data.Hourly.SnowDepth[i]*100,
			data.Hourly.VPD[i],
		)
		sb.WriteString(row)

		if (i+1)%6 == 0 && i < len(data.Hourly.Time)-1 {
			sb.WriteString("\n")
		}
	}
	sb.WriteString("</pre>")

	return sb.String()
}

func buildOpenMeteoURL(lat, lon float64, elevation int) string {
	params := url.Values{}
	params.Set("latitude", fmt.Sprintf("%.4f", lat))
	params.Set("longitude", fmt.Sprintf("%.4f", lon))
	params.Set("timezone", "auto")
	params.Set("forecast_days", "1")
	params.Set("hourly", strings.Join(defaultHourlyParams, ","))

	if elevation > 0 {
		params.Set("elevation", strconv.Itoa(elevation))
	}

	return OpenMeteoBaseURL + "?" + params.Encode()
}

type OpenMeteoResponse struct {
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature   []float64 `json:"temperature_2m"`
		Precipitation []float64 `json:"precipitation"`
		WindSpeed     []float64 `json:"wind_speed_10m"`
		WindDir       []int     `json:"wind_direction_10m"`
		UvIndex       []float64 `json:"uv_index"`
		SnowDepth     []float64 `json:"snow_depth"`
		FreezingLevel []float64 `json:"freezing_level_height"`
		LiftedIndex   []float64 `json:"lifted_index"`
		VPD           []float64 `json:"vapour_pressure_deficit"`
	} `json:"hourly"`
}
