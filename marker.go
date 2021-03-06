package feedly

import "net/url"

// MarkersConuntsResponse : GET /v3/markers/counts
type MarkersConuntsResponse struct {
	Unreadcounts []struct {
		Count   int    `json:"count"`
		Updated int64  `json:"updated"`
		ID      string `json:"id"`
	} `json:"unreadcounts"`
	Updated int64 `json:"updated"`
}

// MarkersCount : https://developer.feedly.com/v3/markers/
func (f *Feedly) MarkersCount(options ...url.Values) (MarkersConuntsResponse, error) {
	result := &MarkersConuntsResponse{}
	option := url.Values{}
	for _, input := range options {
		if err := f.setOption(&option, input); err != nil {
			return *result, err
		}
	}
	_, err := f.request("GET", markerCountURI, result, option)
	return *result, err
}
