package service

import (
	"sync"
	"time"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/slice"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/presenter"
)

type NumberService struct {
	Presenter     *presenter.NumberPresenter
	Fetcher       *fetcher.NumberFetcher
	result        *entity.Numbers
	closedChannel bool
	sync.RWMutex
}

func (n *NumberService) Run(urls []string, ticker *time.Ticker) {
	n.fetchMultiple(urls, ticker)
}

func (n *NumberService) fetchsSingle(url string, ch chan *entity.Numbers) {
	numbers, err := n.Fetcher.Fetch(url)
	if err == nil && !n.isTimesUP() {
		ch <- numbers
	}
}

func (n *NumberService) addResult(num *entity.Numbers) {
	n.Lock()
	n.result.Numbers = slice.SortAndRemoveDuplicatesNumbers(append(n.result.Numbers, num.Numbers...))
	n.Presenter.Parse(n.result)
	n.Unlock()
}

func (n *NumberService) timesUP(up bool) {
	n.Lock()
	defer n.Unlock()
	n.closedChannel = up
}

func (n *NumberService) isTimesUP() bool {
	n.RLock()
	defer n.RUnlock()
	return n.closedChannel
}

func (n *NumberService) setTimeout(ticker *time.Ticker, ch chan *entity.Numbers) {
	for {
		select {
		case _ = <-ticker.C:
			if !n.isTimesUP() {
				ticker.Stop()
				n.timesUP(true)
				close(ch)
				n.Fetcher.Requester.CancelRequest()
			}
		}
	}
}

func (n *NumberService) fetchMultiple(urls []string, ticker *time.Ticker) {
	counter := len(urls)
	ch := make(chan *entity.Numbers, len(urls))

	for _, url := range urls {
		go n.fetchsSingle(url, ch)
	}

	go n.setTimeout(ticker, ch)

	for r := range ch {
		counter--
		n.addResult(r)
		if counter <= 0 && !n.isTimesUP() {
			n.timesUP(true)
			ticker.Stop()
			close(ch)
			break
		}
	}
	n.Presenter.AddHeader("Content-Type", "application/json")
	n.Presenter.Present()
}

func NewNumberService(p *presenter.NumberPresenter, f *fetcher.NumberFetcher) *NumberService {
	return &NumberService{
		Presenter:     p,
		Fetcher:       f,
		result:        &entity.Numbers{Numbers: []int{}},
		closedChannel: false,
	}
}
