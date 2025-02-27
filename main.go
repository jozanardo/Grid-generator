package main

import (
	"fmt"
	"log"

	"grid_generator/internal/reader"
	"grid_generator/internal/scheduling"
	"grid_generator/internal/validation"
	"grid_generator/models"
)

func main() {
	sheetName := "Dados"
	lines, err := reader.ReadExcelFile(sheetName)
	if err != nil {
		log.Fatal("Error reading Excel file:", err)
	}

	var teachers []models.Teacher
	errorsOccurred := false

	for i, line := range lines[1:] {
		teacher, errors := validation.ProcessTeacherData(line)

		if len(errors) > 0 {
			fmt.Printf("⚠️ Errors on line %d:\n", i+2)
			for _, err := range errors {
				fmt.Println("   -", err)
			}
			errorsOccurred = true
		} else {
			teachers = append(teachers, teacher)
		}
	}

	if errorsOccurred {
		fmt.Println("🚨 Errors found. Please correct the data and try again.")
		return
	}

	scheduling.GenerateSchedule(teachers)
}
