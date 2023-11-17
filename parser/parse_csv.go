package parser

import (
	"encoding/csv"
	"os"
)

// this function read the csv from the specified file path
// and return headers,data (records other than headers),error if any
func ParseCSV(filePath string) ([]string, [][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	headers := records[0]
	data := records[1:]

	return headers, data, nil
}
