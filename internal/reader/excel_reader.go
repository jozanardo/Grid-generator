package reader

import (
	"errors"
	"log"

	"github.com/sqweek/dialog"
	"github.com/xuri/excelize/v2"
)

// Opens and reads the specified sheet from an Excel file.
func ReadExcelFile(sheetName string) ([][]string, error) {
	file, err := dialog.File().Filter("Excel Files", "xlsx").Title("Select an Excel file").Load()
	if err != nil {
		log.Fatal("No file selected:", err)
	}

	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	tabs := f.GetSheetList()
	exists := false
	for _, name := range tabs {
		if name == sheetName {
			exists = true
			break
		}
	}
	if !exists {
		return nil, errors.New("sheet not found: " + sheetName)
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, errors.New("failed to read sheet data")
	}

	if len(rows) < 2 {
		return nil, errors.New("spreadsheet is empty or lacks sufficient data")
	}

	return rows, nil
}
