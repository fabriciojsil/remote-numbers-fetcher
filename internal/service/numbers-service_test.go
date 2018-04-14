package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher/numberfetcher"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/request"
)

func TestNumbersService(t *testing.T) {
	t.Run("Fetching one endpoints ", func(t *testing.T) {
		expectedBody := `{"numbers":[2,3,5,7,11,13]}`

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]interface{}{"numbers": []int{2, 3, 5, 7, 11, 13}})
		}))

		writer := httptest.NewRecorder()
		presenter := FakePresenter{Writer: writer}
		requester := request.Requester{Tr: &http.Transport{}}
		fetcher := numberfetcher.NumberFetcher{Requester: requester}

		service := NewNumberService(fetcher, presenter)

		service.Run([]string{server.URL + "?req=1", server.URL + "?req=2"})

		res := writer.Result()
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		bodyString := string(body)

		if !reflect.DeepEqual(bodyString, expectedBody) {
			t.Errorf("Expected %v | Actual %v", expectedBody, bodyString)
		}

	})
}
