package go_sportradar_football

import (
	"time"
)

type Meta struct {
	GeneratedAt time.Time `json:"generated_at"`
	Schema      string    `json:"schema"`
}
