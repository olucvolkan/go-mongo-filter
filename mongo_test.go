package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockRepo struct{}

func (mr *mockRepo) FindRecordsWithCreatedAtAndTotalCounts(params *FindRecordsWithCreatedAtAndTotalCountsParams) ([]Row, error) {
	var result []Row
	result = append(result, Row{ID: "test", Key: "test", CreatedAt: time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC), TotalCount: 10})
	return result, nil
}

func TestMongoHandlerRequest_UnmarshalJSON(t *testing.T) {
	var jsonStr = []byte(`{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}`)
	req, err := http.NewRequest("POST", "/mongo", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(buildMongoHandler(&mockRepo{}))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := 0
	totalCount := 10
	var mongoHandlerResponse MongoHandlerResponse
	json.Unmarshal([]byte(rr.Body.String()), &mongoHandlerResponse)
	if mongoHandlerResponse.Code != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	if mongoHandlerResponse.Records[0].TotalCount != totalCount {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), totalCount)
	}
}
