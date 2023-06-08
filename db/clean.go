package db

import (
	"fmt"
	"time"

	"github.com/Dev-Siri/gedis/models"
	"github.com/Dev-Siri/gedis/utils"
)

func CleanDB() {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		mapJsonContent, err := readFromCache()

		if err != nil {
			fmt.Println("Database cleaner: cleaning iteration failed.")
			return
		}

		for key, data := range mapJsonContent {
			go clearExpiredTTLValues(key, data)
		}
	}
}

func clearExpiredTTLValues(key string, data models.Data) {
	if data.CreatedAt == "" {
		return
	}

	expired, err := utils.IsExpired(data.CreatedAt, data.TTL)

	if err != nil {
		fmt.Printf("Database cleaner: Failed to parse `map[%s].CreatedAt` as datetime\n", key)
		return
	}

	if expired {
		DeleteValue(key)
	}
}