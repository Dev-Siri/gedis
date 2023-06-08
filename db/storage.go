package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Dev-Siri/gedis/models"
)

var storage map[string]models.Data = make(map[string]models.Data)

func InitializeDB() error {
	if err := initCache(); err != nil {
		return err
	}

	mapContent, err := readFromCache()

	if err != nil {
		return err
	}

	storage = mapContent

	return nil
}

func GetValue(key string) models.Data {
	return storage[key]
}

func SetValue(key string, value string, ttl int) {
	createdAt := time.Now().UTC().Format("2006-01-02 15:04:05")

	go writeToCache(key, value, ttl, createdAt)

	if ttl == 0 {
		storage[key] = models.Data{Value: value}
		return
	}

	storage[key] = models.Data{
		Value: value,
		CreatedAt: createdAt,
		TTL: ttl,
	}
}

func DeleteValue(key string) {
	go deleteFromCache(key)
	delete(storage, key)
}

func DeleteAll() {
	emptyMap := make(map[string]models.Data)
	updateCache(emptyMap)
	storage = emptyMap
}

func Increment(key string, amount int) {
	value := GetValue(key)

	valueAsInt, err := strconv.Atoi(value.Value)

	if err != nil {
		fmt.Println("Cannot increment a non-integer value")
		return
	}

	newValue := valueAsInt + amount

	SetValue(key, fmt.Sprint(newValue), value.TTL)
}

func Decrement(key string, amount int) {
	value := GetValue(key)

	valueAsInt, err := strconv.Atoi(value.Value)

	if err != nil {
		fmt.Println("Cannot decrement a non-integer value")
		return
	}

	newValue := valueAsInt - amount

	SetValue(key, fmt.Sprint(newValue), value.TTL)
}

func getStorage() map[string]models.Data {
	return storage
}