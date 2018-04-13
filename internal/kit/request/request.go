package request

import (
	"net/http"
)

// Request defines a Request interface
type Request interface {
	DoRequest(string) (*http.Response, error)
	CancelRequest()
}

//Requester Struct to do Requests to end points
type Requester struct {
	request *http.Request
	Tr      *http.Transport
	http.Client
}

//DoRequest execute a GET request to an specific url
func (r Requester) DoRequest(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	return r.Client.Do(req)
}

func (r Requester) CancelRequest() {
	r.Tr.CancelRequest(r.request)
}
