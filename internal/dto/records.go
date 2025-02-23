package dto

import (
	"test-task/internal/entities"
	"time"
)

type Record struct {
	ID        int       `json:"id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type RecordsResponse struct {
	Records []Record `json:"records"`
}

func RecordFromEntity(e entities.Record) *Record {
	return &Record{
		ID:        e.ID,
		Data:      e.Data,
		CreatedAt: e.CreatedAt,
	}
}

func RecordsFromEntities(entities []entities.Record) []Record {
	records := make([]Record, len(entities))
	for i, rec := range entities {
		records[i] = *RecordFromEntity(rec)
	}
	return records
}
