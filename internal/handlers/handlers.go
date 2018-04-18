package handlers

import (
	"net/http"
	"time"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/request"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/presenter"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/service"
)

var NumberHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	ticker := time.NewTicker(500 * time.Millisecond)
	presenter := &presenter.NumberPresenter{Writer: w}
	requester := &request.Requester{Tr: &http.Transport{}}
	fetcher := &fetcher.NumberFetcher{Requester: requester}
	service := service.NewNumberService(presenter, fetcher)
	urls := r.URL.Query()["u"]
	service.Run(urls, ticker)
})
