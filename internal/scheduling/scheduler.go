package scheduling

import (
	"encoding/json"
	"fmt"
	"log"

	"grid_generator/models"
)

// GenerateSchedule processes the list of teachers and generates a schedule (currently just returns JSON).
func GenerateSchedule(teachers []models.Teacher) {
	jsonData, err := json.MarshalIndent(teachers, "", "  ")
	if err != nil {
		log.Fatal("Error converting to JSON:", err)
	}

	fmt.Println("\n📌 Generated Schedule:")
	fmt.Println(string(jsonData))
}
