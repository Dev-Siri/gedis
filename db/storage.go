package db

var storage map[string]string = make(map[string]string)

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

func GetValue(key string) string {
	return storage[key]
}

func SetValue(key string, value string) {
	go writeToCache(key, value)
	storage[key] = value
}

func DeleteValue(key string) {
	go deleteFromCache(key)
	delete(storage, key)
}

func getStorage() map[string]string {
	return storage
}