package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Dev-Siri/gedis/constants"
	"github.com/Dev-Siri/gedis/models"
)

func initCache() error {	
	if _, err := os.Stat(constants.CacheFolder); os.IsNotExist(err) {		
		if err := os.Mkdir(constants.CacheFolder, 0755); err != nil {
			fmt.Println("Failed to create cache folder")
			return err
		}
	}

	if _, err := os.Stat(constants.MapJsonPath); os.IsNotExist(err) {
		if _, err := os.Create(constants.MapJsonPath); err != nil {
			fmt.Println("Failed to create cache map")
			return err
		}

		os.WriteFile(constants.MapJsonPath, []byte("{}"), 0644)
	}

	return nil
}

func readFromCache() (map[string]models.Data, error) {
	fileContent, err := os.ReadFile(constants.MapJsonPath)

	if err != nil {
		return nil, err
	}

	mapJsonContent := make(map[string]models.Data)

	if err := json.Unmarshal(fileContent, &mapJsonContent); err != nil {
		return nil, err
	}

	return mapJsonContent, nil
}

func updateCache(modifiedMapJson map[string]models.Data) {
	newMapJson, jsonError := json.Marshal(modifiedMapJson)

	if jsonError != nil {
		fmt.Println("Failed to convert new values to JSON")
		return
	}

	if err := os.WriteFile(constants.MapJsonPath, newMapJson, 0644); err != nil {
		fmt.Println("Failed to write new values to cache")
		return
	}
}

func writeToCache(key string, value string, ttl int, createdAt string) {
	mapJsonContent, fileReadError := readFromCache()

	if fileReadError != nil {
		fmt.Println("Failed to read cache")
		return
	}

	creationDate := createdAt

	if ttl == 0 {
		creationDate = ""
	}

	mapJsonContent[key] = models.Data{
		Value: value,
		TTL: ttl,
		CreatedAt: creationDate,
	}

	go updateCache(mapJsonContent)
}

func deleteFromCache(key string) {
	mapJsonContent, fileReadError := readFromCache()

	if fileReadError != nil {
		fmt.Println("Failed to read from cache")
	}

	delete(mapJsonContent, key)

	go updateCache(mapJsonContent)
}