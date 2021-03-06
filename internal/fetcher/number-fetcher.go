package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/request"
)

type NumberFetcher struct {
	Requester request.Request
}

//Fetch retrieves Numbers from an endpoit
func (n NumberFetcher) Fetch(url string) (*entity.Numbers, error) {
	numbers := &entity.Numbers{Numbers: []int{}}
	res, err := n.Requester.DoRequest(url)
	if err != nil {
		return nil, err
	}
	n.unmarshalNumbers(res, numbers)
	return numbers, nil
}

func (n NumberFetcher) unmarshalNumbers(res *http.Response, numbers *entity.Numbers) {
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, numbers)
}
