package dto

import "time"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type Channel struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	YoutubeId   string    `json:"youtube_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
	PublishedAt time.Time `json:"published_at"`
	Active      int       `json:"active"`
	Manager     bool      `json:"manager"`
}
