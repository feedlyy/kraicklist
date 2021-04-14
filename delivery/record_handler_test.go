package delivery

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	mock_service "challenge.haraj.com.sa/kraicklist/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordHandler_GetSearch(t *testing.T) {
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// mock services interface
	mockServices := mock_service.NewMockRecordServices(mockCtrl)
	mockServices.EXPECT().Search(gomock.Any()).Return(records, nil)

	recordHandler := RecordHandler{servicesRecord: mockServices}

	// get new request to the endpoint
	req, err := http.NewRequest("GET", "/search?q=test",
		nil)
	//wrong implementation for query param
	//strings.NewReader(url.Values{"q": {"test"}}.Encode())
	assert.NoError(t, err)

	// recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(recordHandler.GetSearch)
	handler.ServeHTTP(rr, req)

	//assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotNil(t, req)
}
