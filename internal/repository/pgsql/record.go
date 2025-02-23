package pgsql

import (
	"context"
	"gorm.io/gorm"
	"test-task/internal/entities"
	"test-task/internal/repository"
)

type recordRepository struct {
	conn *gorm.DB
}

func NewRecordRepository(conn *gorm.DB) repository.RecordRepository {
	return &recordRepository{conn: conn}
}

func (r *recordRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.conn
}

func (r *recordRepository) GetAllRecords(ctx context.Context, tx *gorm.DB) ([]entities.Record, error) {
	var records []entities.Record

	err := r.getDB(tx).WithContext(ctx).Find(&records).Error
	return records, err
}

func (r *recordRepository) GetRandomRecords(ctx context.Context, tx *gorm.DB) ([]entities.Record, error) {
	var records []entities.Record
	err := r.getDB(tx).WithContext(ctx).
		Order("RANDOM()").
		Limit(10).
		Find(&records).Error
	return records, err
}
