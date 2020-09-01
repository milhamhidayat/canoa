package entity

import "time"

// Team is an entity for team
type Team struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Founded   int32     `json:"founded"`
	Stadium   string    `json:"stadium"`
	CreatedAt time.Time `json:"created_at"`
}
