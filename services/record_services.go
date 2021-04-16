package services

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
	"github.com/sahilm/fuzzy"
	"log"
	"sort"
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

		querys := strings.Split(strings.ToLower(query), " ")
		title := strings.ToLower(record.Title)
		contentCheck := []string{strings.ToLower(record.Content)}
		titleCheck := []string{strings.ToLower(record.Title)}

		checkContent := fuzzy.Find(strings.ToLower(query), contentCheck)
		score := 0
		for _, fuzzys := range checkContent {
			score = fuzzys.Score
		}

		checkTitle := fuzzy.Find(strings.ToLower(query), titleCheck)
		scoreTitles := 0
		for _, fuzzys := range checkTitle {
			scoreTitles = fuzzys.Score
		}

		/*terms:
			1. search by exact title
			2. search by score fuzzy and 'first key word' if it contains more than 1 word
		       ex: iphone black, then the keyword are iphone*/

		if strings.Contains(title, strings.ToLower(query)) {
			s.recordScore = append(s.recordScore, entity.RecordWithScore{
				Record: record,
				Score:  int64(scoreTitles),
			})
		} else if score >= 0 && strings.Contains(title, querys[0]) {
			s.recordScore = append(s.recordScore, entity.RecordWithScore{
				Record: record,
				Score:  int64(score),
			})
		}
	}

	// Sort by it's score
	sort.SliceStable(s.recordScore, func(i, j int) bool {
		return s.recordScore[i].Score > s.recordScore[j].Score
	})

	//loop the record that already sorted
	for _, value := range s.recordScore {
		result = append(result, value.Record)
	}

	cached[strings.ToLower(query)] = result
	check[strings.ToLower(query)] = true

	//reset the struct
	s.recordScore = nil

	return result, nil
}
