package configs

import (
	"os"
	"strconv"
)

func GetCacheConfig() (string, string, string, int) {
	host := os.Getenv("CACHE_HOST")
	if len(host) <= 0 {
		host = "0.0.0.0"
	}

	port := os.Getenv("CACHE_PORT")
	if len(port) <= 0 {
		port = "6379"
	}

	passwd := os.Getenv("CACHE_PASSWORD")
	dbString := os.Getenv("CACHE_DB")
	if len(dbString) <= 0 {
		dbString = "0"
	}
	db, err := strconv.Atoi(dbString)
	if err != nil {
		panic(err)
	}

	return host, port, passwd, db
}
