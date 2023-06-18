package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Dev-Siri/gedis/constants"
	"github.com/Dev-Siri/gedis/models"
	"github.com/Dev-Siri/gedis/utils"
)

func MakeBackup(backupType string) {
	if _, err := os.Stat(constants.BackupFolder); os.IsNotExist(err) {		
		if err := os.Mkdir(constants.BackupFolder, 0755); err != nil {
			fmt.Println("(Error) Failed to create backup folder")
			return
		}
	}

	storage := getStorage()
	backupFilePath := constants.BackupFolder + "/backup-" + utils.GetCurrentTimestamp() + "." + backupType

	if _, err := os.Stat(backupFilePath); os.IsNotExist(err) {
		if _, err := os.Create(backupFilePath); err != nil {
			fmt.Println("(Error) Failed to create backup snapshot file")
			return
		}

		if backupType == "csv" {
			backupFile, err := os.Create(backupFilePath)

			if err != nil {
				fmt.Println("(Error) Failed to create CSV backup file")
				return
			}
	
			if err := utils.EncodeToCSV(storage, backupFile); err != nil {
				fmt.Println("(Error) Failed to encode backup as CSV")
				return
			}
		} else {
			backupContent, jsonError := json.Marshal(storage)
	
			if jsonError != nil {
				fmt.Println("(Error) Failed to encode backup as JSON")
				return
			}
	
			os.WriteFile(backupFilePath, backupContent, 0644)			
		}
	}
}

func LoadBackup(filename string) {
	filePath := "backups/" + filename
	fileType := strings.Split(filename, ".")[1]

	if fileType == "csv" {
		file, csvOpenError := os.Open(filePath)

		if csvOpenError != nil {
			fmt.Println("(Error) Failed to read backup as CSV")
			return
		}

		defer file.Close()

		decodedMap, decodeError := utils.DecodeCSV(file)

		if decodeError != nil {
			fmt.Println("(Error) Failed to decode backup")
			return
		}

		for key, data := range decodedMap {
			SetValue(key, data.Value, data.TTL)
		}
	} else if fileType == "json" {
		backupBytes, backupReadError := os.ReadFile(filePath)

		if backupReadError != nil {
			fmt.Println("(Error) Failed to read backup as JSON")
			return
		}

		data := make(map[string]models.Data)

		if err := json.Unmarshal(backupBytes, &data); err != nil {
			fmt.Println("(Error) Failed to decode backup")
			return
		}

		for key, data := range data {
			SetValue(key, data.Value, data.TTL)
		}
	} else {
		fmt.Println("(Error) Unknown backup file type")
		return
	}

	fmt.Println("(Info) Backup loaded into memory")
}