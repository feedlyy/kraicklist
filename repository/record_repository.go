package repository

import (
	"bufio"
	"challenge.haraj.com.sa/kraicklist/entity"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
)

type RecordRepository interface {
	Load (filepath string) (*entity.Searcher, error)
}

type RecordRepo struct {
	searcher *entity.Searcher
}

func NewRecordRepository(searcher entity.Searcher) *RecordRepo {
	return &RecordRepo{searcher: &searcher}
}

func (repo *RecordRepo) Load(filepath string) (*entity.Searcher, error) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()
	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}
	// read the reader using scanner to construct records
	var records []entity.Record
	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		r := entity.Record{}
		err = json.Unmarshal(cs.Bytes(), &r)
		if err != nil {
			continue
		}
		records = append(records, r)
	}
	repo.searcher.Records = records

	return repo.searcher, nil
}
