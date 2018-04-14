package presenter

import "github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"

type Presenter interface {
	Present(entity.Numbers)
}
