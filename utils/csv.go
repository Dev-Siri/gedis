package utils

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/Dev-Siri/gedis/models"
)

func EncodeToCSV(data map[string]models.Data, file *os.File) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Key", "Value"}

	if err := writer.Write(header); err != nil {
		return err
	}

	for key, data := range data {
		row := []string{key, data.Value}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	
	return nil
}

func DecodeCSV(file *os.File) (map[string]models.Data, error) {
	reader := csv.NewReader(file)

	data := make(map[string]models.Data)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if len(row) == 2 {
			key := row[0]
			value := row[1]
			data[key] = models.Data{Value: value}
		}
	}

	delete(data, "Key")

	return data, nil
}