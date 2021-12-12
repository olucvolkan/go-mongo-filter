package main

import (
	"encoding/json"
	"fmt"
	http "net/http"
	"sync"
	"time"
)

const (
	SuccessStatus int = 0
	FailStatus    int = 1
)

type MongoHandlerRequest struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	MinCount  int       `json:"minCount"`
	MaxCount  int       `json:"maxCount"`
}

type InMemoryRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MongoHandlerResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Records []record `json:"records"`
}

type InMemoryResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}

// Record
type record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type StoreKeyValuePairRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InMemoryMap struct {
	KeyValuePair map[string]string
	Mutex        *sync.Mutex
}

var customMap *InMemoryMap

func (r *MongoHandlerRequest) UnmarshalJSON(j []byte) error {
	var raw map[string]interface{}

	err := json.Unmarshal(j, &raw)
	if err != nil {
		return err
	}

	dateFormat := "2006-01-02"

	startDate, err := time.Parse(dateFormat, raw["startDate"].(string))

	if err != nil {
		return err
	}

	r.StartDate = startDate

	endDate, err := time.Parse(dateFormat, raw["endDate"].(string))

	if err != nil {
		return err
	}

	r.EndDate = endDate

	minCount := raw["minCount"].(float64)
	r.MinCount = int(minCount)

	maxCount := raw["maxCount"].(float64)
	r.MaxCount = int(maxCount)
	return nil
}

func buildMongoHandler(repo Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&MongoHandlerResponse{
				Code:    FailStatus,
				Message: "Not Found",
				Records: nil,
			})
			return
		}

		// Decode the post body
		var request MongoHandlerRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code": FailStatus,
				"msg":  "Invalid request",
			})
			return
		}

		// Query MongoDB via repo
		records, err := repo.FindRecordsWithCreatedAtAndTotalCounts(&FindRecordsWithCreatedAtAndTotalCountsParams{
			CreatedAtAfter:  request.StartDate,
			CreatedAtBefore: request.EndDate,
			TotalCountsFrom: request.MinCount,
			TotalCountsTo:   request.MaxCount,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&MongoHandlerResponse{
				Code:    FailStatus,
				Message: err.Error(),
				Records: nil,
			})
			return
		}

		// Encode results
		json.NewEncoder(w).Encode(&MongoHandlerResponse{
			Code:    SuccessStatus,
			Message: "success",
			Records: formatRows(records),
		})
	}
}

func buildInMemoryHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			InMemoryPostHandler(kvstore)(w, r)
		} else if r.Method == "GET" {
			InMemoryGetHandler(kvstore)(w, r)
		}
	}
}

func InMemoryPostHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request InMemoryRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&InMemoryResponse{Code: FailStatus, Message: "invalid request"})
			return
		}

		if ok, err := kvstore.Set(request.Key, request.Value); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&InMemoryResponse{Code: FailStatus, Message: err.Error()})
			return
		}

		// Encode results
		json.NewEncoder(w).Encode(&InMemoryResponse{
			Key:   request.Key,
			Value: request.Value,
		})
	}
}

func InMemoryGetHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code": FailStatus,
				"msg":  "Invalid request",
			})
			return
		}

		key := r.URL.Query().Get("key")
		value, err := kvstore.Get(key)

		if err != nil {
			return
		}
		// Encode results
		json.NewEncoder(w).Encode(&InMemoryResponse{
			Key:   key,
			Value: value,
		})
	}
}

// Set : In-memory rest api set endpoint
func Set(key string, val interface{}) {
	customMap.Mutex.Lock()
	customMap.KeyValuePair[key] = fmt.Sprint(val)
	customMap.Mutex.Unlock()
}

func Get(key string) string {
	customMap.Mutex.Lock()
	defer customMap.Mutex.Unlock()
	if val, ok := customMap.KeyValuePair[key]; ok {
		return val
	}
	return ""
}

func formatRows(rows []Row) []record {
	result := make([]record, 0)

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		result = append(result, record{
			Key:        row.Key,
			CreatedAt:  row.CreatedAt.Local(),
			TotalCount: row.TotalCount,
		})
	}
	return result
}
