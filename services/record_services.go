package services

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
	"github.com/sahilm/fuzzy"
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
	recordScore []entity.RecordWithScore
}

type Recordss []entity.Record

func NewRecordServices(record repository.RecordRepository) *RecordServ {
	return &RecordServ{repoRecord: record}
}

func (e Recordss) String(i int) string {
	if e[i].Title > e[i].Content {
		return e[i].Title
	} else if e[i].Content > e[i].Title {
		return e[i].Content
	}
	return e[i].Title + e[i].Content
}

func (e Recordss) Len() int {
	return len(e)
}

func (s *RecordServ) Search(query string) ([]entity.Record, error) {
	var result []entity.Record
	records, err := s.repoRecord.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	fill := append(Recordss{}, records.Records...)

	finds := fuzzy.FindFrom(strings.ToLower(query), fill)
	for _, r := range finds {
		if r.Score > 0 && strings.Contains(strings.ToLower(r.Str), strings.ToLower(query)) {
			result = append(result, fill[r.Index])
		} else if r.Score > 0 {
			result = append(result, fill[r.Index])
		}
	}

	cached[strings.ToLower(query)] = result
	check[strings.ToLower(query)] = true

	//reset the struct
	s.recordScore = nil

	return result, nil
}
