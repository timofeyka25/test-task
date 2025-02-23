package entities

import "time"

type Record struct {
	ID        int
	Data      string
	CreatedAt time.Time
}
