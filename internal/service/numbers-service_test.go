package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/request"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/presenter"
)

func TestNumbersService(t *testing.T) {
	t.Run("Fetching two success endpoints ", func(t *testing.T) {
		expectedBody := `{"numbers":[2,3,4,6,7,9,11,13,15]}`
		server := createServer()

		writer := httptest.NewRecorder()
		presenter := &presenter.NumberPresenter{Writer: writer}
		requester := &request.Requester{Tr: &http.Transport{}}
		fetcher := &fetcher.NumberFetcher{Requester: requester}

		service := NewNumberService(presenter, fetcher)
		ticker := time.NewTicker(time.Millisecond * 4)
		service.Run([]string{server.URL + "?req=1", server.URL + "?req=2"}, ticker)

		res := writer.Result()
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		bodyString := string(body)

		if !reflect.DeepEqual(bodyString, expectedBody) {
			t.Errorf("Expected %v | Actual %v", expectedBody, bodyString)
		}
	})

	t.Run("Fetching one success and one error endpoints ", func(t *testing.T) {
		expectedBody := `{"numbers":[2,3,4,7,11,13,15]}`
		server := createServer()

		writer := httptest.NewRecorder()
		presenter := &presenter.NumberPresenter{Writer: writer}
		requester := &request.Requester{Tr: &http.Transport{}}
		fetcher := &fetcher.NumberFetcher{Requester: requester}

		service := NewNumberService(presenter, fetcher)
		ticker := time.NewTicker(time.Millisecond * 3)
		service.Run([]string{server.URL + "?req=1", server.URL + "?req=4"}, ticker)

		res := writer.Result()
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		bodyString := string(body)

		if !reflect.DeepEqual(bodyString, expectedBody) {
			t.Errorf("Expected %v | Actual %v", expectedBody, bodyString)
		}
	})

	t.Run("Fetching one success and one later endpoints ", func(t *testing.T) {
		expectedBody := `{"numbers":[2,3,4,7,11,13,15]}`
		server := createServer()

		writer := httptest.NewRecorder()
		presenter := &presenter.NumberPresenter{Writer: writer}
		requester := &request.Requester{Tr: &http.Transport{}}
		fetcher := &fetcher.NumberFetcher{Requester: requester}

		service := NewNumberService(presenter, fetcher)

		ticker := time.NewTicker(time.Millisecond * 3)
		service.Run([]string{server.URL + "?req=1", server.URL + "?req=3"}, ticker)
		res := writer.Result()
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		bodyString := string(body)

		if !reflect.DeepEqual(bodyString, expectedBody) {
			t.Errorf("Expected %v | Actual %v", expectedBody, bodyString)
		}
	})

}

func createServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := r.URL.Query().Get("req")
		switch req {
		case "1":
			json.NewEncoder(w).Encode(map[string]interface{}{"numbers": []int{2, 3, 7, 4, 11, 13, 15}})
		case "2":
			json.NewEncoder(w).Encode(map[string]interface{}{"numbers": []int{3, 2, 7, 11, 6, 9}})
		case "3":
			time.Sleep(time.Duration(5 * time.Millisecond))
			json.NewEncoder(w).Encode(map[string]interface{}{"numbers": []int{3, 2, 7, 11, 6, 9}})
		case "4":
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		}
	}))
	return server
}
