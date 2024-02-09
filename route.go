package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("channel_id")

	if channelID == "" {
		http.Error(w, "channel_id parameter is missing", http.StatusBadRequest)
		return
	}

	cacheItem, hit := cache[channelID]
	if hit && time.Now().Before(cacheItem.ExpiresAt) {

		response := Response{
			ChannelURL: fmt.Sprintf("https://www.youtube.com/channel/%s", cacheItem.ChannelID),
			VideoURL:   fmt.Sprintf("https://www.youtube.com/watch?v=%s", cacheItem.VideoID),
		}

		data, err := json.Marshal(response)

		if err != nil {
			log.Fatalf("Failed to read cache: %v", err)
			http.Error(w, "Failed to read cache", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)

		log.Default().Printf("Cache hit")
		return
	}

	call := service.Search.List([]string{"snippet"}).
		ChannelId(channelID).
		MaxResults(10).
		Order("date").
		Type("video")

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
		http.Error(w, "Failed to fetch YouTube", http.StatusInternalServerError)
	}

	found := false
	for _, item := range response.Items {
		if found {
			break
		}

		if strings.Contains(item.Snippet.Title, "#shorts") {
			continue
		}

		found = true

		cacheItem = CacheItem{
			ChannelID: channelID,
			VideoID:   item.Id.VideoId,
			ExpiresAt: time.Now().Add(1 * time.Hour),
		}

		mutex.Lock()
		cache[channelID] = cacheItem
		mutex.Unlock()

		response := Response{
			ChannelURL: fmt.Sprintf("https://www.youtube.com/channel/%s", cacheItem.ChannelID),
			VideoURL:   fmt.Sprintf("https://www.youtube.com/watch?v=%s", cacheItem.VideoID),
		}

		data, err := json.Marshal(response)

		if err != nil {
			log.Fatalf("Failed to read cache: %v", err)
			http.Error(w, "Failed to read cache", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}
