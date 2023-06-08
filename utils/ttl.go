package utils

import "time"

func IsExpired(creationDate string, ttl int) (bool, error) {
	createdAt, err := time.ParseInLocation("2006-01-02 15:04:05", creationDate, time.UTC) 

	if err != nil {
		return false, err
	}

	expirationTime := createdAt.Add(time.Duration(ttl) * time.Second)
	currentTime := time.Now()

	return currentTime.After(expirationTime), nil
}