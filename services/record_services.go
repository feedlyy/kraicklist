package services

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
	"log"
	"strings"
)

var path = "data.gz"

type RecordServices interface {
	Search(query string) ([]entity.Record, error)
}

type RecordServ struct {
	repoRecord repository.RecordRepository
}

func NewRecordServices(record repository.RecordRepository) *RecordServ {
	return &RecordServ{repoRecord: record}
}

func (s *RecordServ) Search(query string) ([]entity.Record, error) {
	var result []entity.Record
	records, err := s.repoRecord.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records.Records {
		if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) {
			result = append(result, record)
		}
	}
	return result, nil
}
