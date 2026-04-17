package tests

import (
	"strings"
	"testing"

	"goweatherbot/handlers"
)

func TestFormatWeatherResponse(t *testing.T) {
	// 1. Готовим тестовые данные
	mockData := handlers.OpenMeteoResponse{}
	mockData.Hourly.Time = []string{"2026-01-03T12:00"}
	mockData.Hourly.Temperature = []float64{-5.5}
	mockData.Hourly.Precipitation = []float64{1.2}
	mockData.Hourly.WindSpeed = []float64{18.0} // 18 км/ч / 3.6 = 5.0 м/с
	mockData.Hourly.WindDir = []int{135}        // 135 градусов = ЮВ
	mockData.Hourly.SnowDepth = []float64{0.25} // 0.25м = 25см
	mockData.Hourly.VPD = []float64{0.15}

	locationName := "Роза Пик"

	// 2. Вызываем функцию
	result := handlers.FormatWeatherResponse(mockData, locationName)

	// 3. Проверки (Assertions)

	// Проверка заголовка
	if !strings.Contains(result, "Локация: Роза Пик") {
		t.Errorf("Заголовок локации не найден или неверен")
	}

	// Проверка конвертации ветра (18.0 км/ч -> 5.0 м/с)
	if !strings.Contains(result, " 5.0") {
		t.Errorf("Скорость ветра не конвертировалась в 5.0 м/с. Результат: %s", result)
	}

	// Проверка направления ветра (135 градусов -> ЮВ)
	if !strings.Contains(result, "ЮВ") {
		t.Errorf("Направление ветра для 135 градусов должно быть ЮВ")
	}

	// Проверка снега (0.25м -> 25см)
	if !strings.Contains(result, "  25") {
		t.Errorf("Снег не конвертировался в 25см")
	}

	// Проверка VPD
	if !strings.Contains(result, " 0.15") {
		t.Errorf("Значение VPD не найдено в таблице")
	}

	// Проверка наличия HTML-тегов
	if !strings.HasPrefix(result, "<b>") || !strings.Contains(result, "<pre>") {
		t.Errorf("Отсутствуют обязательные HTML теги форматирования")
	}
}

func TestEmptyData(t *testing.T) {
	mockData := handlers.OpenMeteoResponse{}
	result := handlers.FormatWeatherResponse(mockData, "Тест")

	expected := "Нет данных для отображения"
	if result != expected {
		t.Errorf("Ожидалось '%s', получено '%s'", expected, result)
	}
}
