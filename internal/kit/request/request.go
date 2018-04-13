package request

import (
	"net/http"
)

// Request defines a Request interface
type Request interface {
	DoRequest(string) (*http.Response, error)
}

//Requester Struct to do Requests to end points
type Requester struct {
	Client http.Client
}

//DoRequest execute a GET request to an specific url
func (r Requester) DoRequest(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	return r.Client.Do(req)
}
