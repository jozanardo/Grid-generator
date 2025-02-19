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
		teacher, err := validation.ValidateTeacher(line)
		if err != nil {
			errorsOccurred = true
			fmt.Printf("⚠️ Error on line %d: %v\n", i+2, err)
		} else {
			teachers = append(teachers, *teacher)
		}
	}

	if errorsOccurred {
		fmt.Println("🚨 Errors found. Please correct the data and try again.")
		return
	}

	// Call scheduler to generate and display the schedule
	scheduling.GenerateSchedule(teachers)
}
