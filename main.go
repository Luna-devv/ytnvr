package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	_ "github.com/joho/godotenv/autoload"
)

var apiKey = getEnvValue("API_KEY")

var cache = make(map[string]CacheItem)
var mutex sync.Mutex

var service, err = youtube.NewService(context.Background(), option.WithAPIKey(apiKey))

func main() {

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", fetchHandler)
	fmt.Println("Server is running on :8080")

	http.ListenAndServe(":8080", nil)

}

func getEnvValue(key string) string {
	value, set := os.LookupEnv(key)
	if !set {
		log.Fatalf("Config variable %s was missing\n", key)
	}
	return value
}
