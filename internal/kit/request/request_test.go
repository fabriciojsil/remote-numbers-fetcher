package request

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {

		expectedBody := "body response"
		fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(expectedBody))
		})

		requester, url := createRequest(fakeHandler)
		response, err := requester.DoRequest(url)

		if err != nil {
			t.Errorf(" It shouldn 't thrown an error %s ", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != expectedBody {
			t.Errorf("Expected %s Actual %s", expectedBody, string(body))
		}
	})

	t.Run("Sevice unavailable", func(t *testing.T) {
		fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		})
		requester, url := createRequest(fakeHandler)
		response, err := requester.DoRequest(url)

		if err != nil {
			t.Errorf("It mustn't throw an error %s ", err)
		}

		if response.StatusCode != http.StatusServiceUnavailable {
			t.Errorf("Expected %v Actual %v", http.StatusServiceUnavailable, response.StatusCode)
		}
	})

	t.Run("Nonexistent URL", func(t *testing.T) {

		client := http.Client{Transport: &http.Transport{}}
		requester := Requester{Client: client}
		_, err := requester.DoRequest("https://localhost/inexistent")

		if err == nil {
			t.Errorf("It must thrown an error %s ", err)
		}
	})

	t.Run("Invalid URL", func(t *testing.T) {

		client := http.Client{Transport: &http.Transport{}}
		requester := Requester{Client: client}
		_, err := requester.DoRequest("localhost/")

		if err == nil {
			t.Errorf("It must thrown an error %s ", err)
		}
	})
}

func createRequest(handler http.HandlerFunc) (Requester, string) {
	w := httptest.NewServer(handler)
	client := http.Client{Transport: &http.Transport{}}

	return Requester{Client: client}, w.URL
}
