package numberservice

import (
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/fetcher"
	"github.com/fabriciojsil/remote-numbers-fetcher/internal/presenter"
)

type NumberService struct {
	Fetcher   fetcher.Fetcher
	Presenter presenter.Presenter
}

func (n NumberService) Run(urls []string) {
	numbers := n.Fetcher.Fetch(urls[0])
	n.Presenter.Present(numbers)
}

func newNumberService(f fetcher.Fetcher, p presenter.Presenter) NumberService {
	return NumberService{
		Fetcher:   f,
		Presenter: p,
	}
}
