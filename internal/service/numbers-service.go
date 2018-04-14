package service

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/kit/slice"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/presenter"
)

type numberService struct {
	Fetcher   fetcher.Fetcher
	Presenter presenter.Presenter
	result    entity.Numbers
	sync.Mutex
}

func (n *numberService) Run(urls []string) {
	n.fetchMultiple(urls)
	n.Presenter.Present(n.result)
}

func (n *numberService) addResult(num *entity.Numbers) {
	n.Lock()
	n.result.Numbers = slice.OrderToRemoveDuplicates(append(n.result.Numbers, num.Numbers...))
	n.Unlock()
}

func (n *numberService) fetchInRoutines(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	n.addResult(n.Fetcher.Fetch(url))
}

func (n *numberService) fetchMultiple(urls []string) {
	waitGroup := new(sync.WaitGroup)
	in := make(chan string, len(urls))

	for _, url := range urls {
		waitGroup.Add(1)
		go n.fetchInRoutines(url, waitGroup)
		in <- url
	}
	close(in)
	waitGroup.Wait()
}

func NewNumberService(f fetcher.Fetcher, p presenter.Presenter) numberService {
	return numberService{
		Fetcher:   f,
		Presenter: p,
		result:    entity.Numbers{Numbers: []int{}},
	}
}

type FakePresenter struct {
	Writer http.ResponseWriter
}

func (f FakePresenter) Present(numbers entity.Numbers) {
	f.Writer.Header().Add("Content-Type", "application/json")
	res, _ := json.Marshal(numbers)
	f.Writer.Write(res)
}
