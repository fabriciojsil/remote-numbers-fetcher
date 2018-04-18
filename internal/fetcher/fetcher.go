package fetcher

import "github.com/fabriciojsil/remote-numbers-fetcher/internal/entity"

//Fetcher Defines a pattern to Fetche Structs
type Fetcher interface {
	Fetch(url string) *entity.Numbers
}
