package handlers

type Location struct {
	Code      string
	Name      string
	Latitude  float64
	Longitude float64
	Elevation int
}

var Locations = map[string]Location{
	"rosapeak": {
		Code:      "rosapeak",
		Name:      "Роза Пик",
		Latitude:  43.6252,
		Longitude: 40.3101,
		Elevation: 2320,
	},

	"rosa1600": {
		Code:      "rosa1600",
		Name:      "Роза 1600",
		Latitude:  43.6388,
		Longitude: 40.3136,
		Elevation: 1600,
	},

	"caucaseexpress": {
		Code:      "caucaseexpress",
		Name:      "Кавказский экспресс",
		Latitude:  43.6445,
		Longitude: 40.3151,
		Elevation: 1350,
	},

	"krokus": {
		Code:      "krokus",
		Name:      "Крокус",
		Latitude:  43.6121,
		Longitude: 40.3314,
		Elevation: 2509,
	},

	"edelweiss": {
		Code:      "edelweiss",
		Name:      "Эдельвейс",
		Latitude:  43.6100,
		Longitude: 40.2862,
		Elevation: 1472,
	},
}
