package fetcher

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/request"
)

func TestNumberFetcher(t *testing.T) {
	t.Run("Sucess request", func(t *testing.T) {

		numbersExpecteds := &entity.Numbers{
			Numbers: []int{2, 3, 5, 7, 11, 13},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]interface{}{"numbers": []int{2, 3, 5, 7, 11, 13}})
		}))

		fetcher := NumberFetcher{
			Requester: request.Requester{Tr: &http.Transport{}},
		}

		numbers, _ := fetcher.Fetch(server.URL)

		if !reflect.DeepEqual(numbersExpecteds, numbers) {
			t.Errorf("Expected %v | Actual %v", numbersExpecteds, numbers)
		}

	})

	t.Run("Returning an error on request", func(t *testing.T) {
		expectedError := errors.New("New error")
		fetcher := NumberFetcher{
			Requester: fakeRequest{},
		}

		_, err := fetcher.Fetch("doesnt matters")

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("Expected %v | Actual %v", err, expectedError)
		}
	})

}

type fakeRequest struct {
}

func (f fakeRequest) DoRequest(string) (*http.Response, error) {
	return nil, errors.New("New error")
}
func (f fakeRequest) CancelRequest() {}
