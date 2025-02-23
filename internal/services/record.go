package services

import (
	"context"
	"errors"
	"test-task/internal/dto"
	"test-task/internal/repository"
)

type RecordService struct {
	recordRepo repository.RecordRepository
}

func NewRecordService(recordRepo repository.RecordRepository) *RecordService {
	return &RecordService{
		recordRepo: recordRepo,
	}
}

func (s *RecordService) GetAllRecords(ctx context.Context) (*dto.RecordsResponse, error) {
	records, err := s.recordRepo.GetAllRecords(ctx, nil)
	if err != nil {
		return nil, errors.New("failed to get all records")
	}

	return &dto.RecordsResponse{
		Records: dto.RecordsFromEntities(records),
	}, nil
}
