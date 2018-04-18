package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher"
)

type NumberPresenter struct {
	Writer  http.ResponseWriter
	Fetcher fetcher.NumberFetcher
	parsed  []byte
}

func (f *NumberPresenter) Parse(numbers *entity.Numbers) {
	res, _ := json.Marshal(numbers)
	f.parsed = res
}
func (f *NumberPresenter) Present() {
	f.Writer.Write(f.parsed)
}

func (f *NumberPresenter) AddHeader(key, value string) {
	f.Writer.Header().Add(key, value)
}
