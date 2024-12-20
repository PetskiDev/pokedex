package main
  
type LocationResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []Location `json:"results"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
