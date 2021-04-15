package services

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
	"log"
	"strings"
)

var path = "data.gz"
var check = make(map[string]bool, 0)
var cached = make(map[string][]entity.Record, 0)

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
		if check[strings.ToLower(query)] {
			log.Println("Cached successful")
			return cached[strings.ToLower(query)], nil
		}

		if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) {
			result = append(result, record)
		}
	}
	cached[strings.ToLower(query)] = result
	check[strings.ToLower(query)] = true
	return result, nil
}
