package request

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

		requester := Requester{Tr: &http.Transport{}}
		_, err := requester.DoRequest("https://localhost/inexistent")

		if err == nil {
			t.Errorf("It must thrown an error %s ", err)
		}
	})

	t.Run("Invalid URL", func(t *testing.T) {

		requester := Requester{Tr: &http.Transport{}}
		_, err := requester.DoRequest("localhost/")

		if err == nil {
			t.Errorf("It must thrown an error %s ", err)
		}
	})

	t.Run("Canceling a reqest", func(t *testing.T) {
		wasCalled := false
		fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(10) * time.Millisecond)
			wasCalled = true
			w.Write([]byte("body response"))
		})

		requester, url := createRequest(fakeHandler)

		go requester.DoRequest(url)
		requester.CancelRequest()
		if wasCalled == true {
			t.Errorf("It doenst be caller %v ", wasCalled)
		}
	})

}

func createRequest(handler http.HandlerFunc) (Requester, string) {
	w := httptest.NewServer(handler)
	return Requester{Tr: &http.Transport{}}, w.URL
}
