package validation

import (
	"fmt"
	"strconv"
	"strings"

	"grid_generator/models"
	"grid_generator/utils"
)

// Validates and normalizes a teacher's data extracted from Excel
func ProcessTeacherData(line []string) (models.Teacher, []string) {
	var errors []string

	teacher := models.Teacher{}

	if len(line) < 5 {
		return teacher, []string{fmt.Sprintf("Incomplete line, expected 5 columns, found: %d", len(line))}
	}

	teacher.Name = strings.TrimSpace(line[0])
	if teacher.Name == "" || !utils.NameRegex.MatchString(teacher.Name) {
		errors = append(errors, fmt.Sprintf("Invalid name: '%s'", teacher.Name))
	}

	teacher.Subject = strings.TrimSpace(line[1])
	normalizedSubject := utils.NormalizeText(teacher.Subject)
	if !utils.IsValidSubject(normalizedSubject) {
		errors = append(errors, fmt.Sprintf("Invalid subject for ['%s']: '%s'", teacher.Name, teacher.Subject))
	}

	numberOfClasses, err := strconv.Atoi(strings.TrimSpace(line[2]))
	if err != nil || numberOfClasses <= 0 {
		errors = append(errors, fmt.Sprintf("Invalid number of classes fo['%s']: '%s'", teacher.Name, line[2]))
	} else {
		teacher.NumberOfClasses = numberOfClasses
	}

	availableDays, dayErrors := ProcessAvailableDays(strings.TrimSpace(line[3]))
	if len(dayErrors) > 0 {
		for _, err := range dayErrors {
			errors = append(errors, fmt.Sprintf("Error in validating days for ['%s']: %v", teacher.Name, err))
		}
	} else {
		teacher.AvailableDays = availableDays
	}

	teacher.AvailableHours = strings.TrimSpace(line[4])

	return teacher, errors
}

// validates and normalizes the days reported to the teacher.
func ProcessAvailableDays(daysInput string) (string, []string) {
	days := strings.Split(daysInput, ",")
	var normalizedDays []string
	var errors []string

	for _, day := range days {
		day = strings.TrimSpace(strings.ToUpper(day))

		if normalized, exists := utils.WeekDayNormalization[day]; exists {
			day = normalized
		}

		if !utils.ValidWeekDays[day] {
			errors = append(errors, fmt.Sprintf("invalid day: '%s'", day))
			continue
		}

		normalizedDays = append(normalizedDays, day)
	}

	return strings.Join(normalizedDays, ","), errors
}
