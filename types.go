package main

import "time"

type CacheItem struct {
	ChannelID string    `json:"channelId"`
	VideoID   string    `json:"videoId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Response struct {
	ChannelURL string `json:"channelUrl"`
	VideoURL   string `json:"videoUrl"`
}
