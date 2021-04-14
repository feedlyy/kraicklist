package services

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	mock_repository "challenge.haraj.com.sa/kraicklist/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordServ_Search(t *testing.T) {
	var records []entity.Record
	for i := 0; i < 2; i++ {
		records = append(records, entity.Record{
			ID: 	   int64(i),
			Title:     "test",
			Content:   "test",
			ThumbURL:  "test",
			Tags:      nil,
			UpdatedAt: 0,
			ImageURLs: nil,
		})
	}
	searches := &entity.Searcher{Records: records}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//mock repo interface
	mockRepo := mock_repository.NewMockRecordRepository(mockCtrl)
	mockRepo.EXPECT().Load(gomock.Any()).Return(searches, nil)

	recordServ := RecordServ{repoRecord: mockRepo}

	res, err := recordServ.Search("test")
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
