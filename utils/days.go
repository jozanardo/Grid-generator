package utils

// Represents the allowed weekdays
var ValidWeekDays = map[string]bool{
	"SEG": true, "TER": true, "QUA": true, "QUI": true, "SEX": true, "SAB": true,
}

// Maps full names to standardized abbreviations
var WeekDayNormalization = map[string]string{
	"SEGUNDA": "SEG",
	"TERÇA":   "TER",
	"QUARTA":  "QUA",
	"QUINTA":  "QUI",
	"SEXTA":   "SEX",
	"SÁBADO":  "SAB",
	"SABADO":  "SAB",
}
