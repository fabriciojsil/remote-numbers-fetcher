package fetcher

//Fetcher Defines a pattern to Fetche Structs
type Fetcher interface {
	Fetch(url string)
}
