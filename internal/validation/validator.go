package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"grid_generator/models"
	"grid_generator/utils"
)

// ValidateTeacher validates a teacher's data from a line in the spreadsheet.
func ValidateTeacher(line []string) (*models.Teacher, error) {
	if len(line) < 5 {
		return nil, errors.New("incomplete row")
	}

	// Validate name
	name := strings.TrimSpace(line[0])
	if name == "" || !utils.NameRegex.MatchString(name) {
		return nil, fmt.Errorf("invalid name: '%s'", name)
	}

	// Validate subject
	subject := strings.TrimSpace(line[1])
	normalizedSubject := utils.NormalizeText(subject) // Remove accents, convert to lowercase
	if !utils.IsValidSubject(normalizedSubject) {
		return nil, fmt.Errorf("invalid subject for '%s': '%s'", name, subject)
	}

	// Validate number of classes
	numberOfClasses, err := strconv.Atoi(strings.TrimSpace(line[2]))
	if err != nil || numberOfClasses <= 0 {
		return nil, fmt.Errorf("invalid number of classes for '%s': '%s'", name, line[2])
	}

	// Trim available days and hours
	availableDays := strings.TrimSpace(line[3])
	availableHours := strings.TrimSpace(line[4])

	return &models.Teacher{
		Name:            name,
		Subject:         subject,
		NumberOfClasses: numberOfClasses,
		AvailableDays:   availableDays,
		AvailableHours:  availableHours,
	}, nil
}
