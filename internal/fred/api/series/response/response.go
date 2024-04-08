package response

import (
	"errors"
	"strconv"
)

// Res represents HTTP response from FRED API.
type Res struct {
	Units        string        `json:"units,omitempty"`
	OutputType   int           `json:"output_type,omitempty"`
	FileType     string        `json:"file_type,omitempty"`
	OrderBy      string        `json:"order_by,omitempty"`
	SortBy       string        `json:"sort_by,omitempty"`
	Count        int           `json:"count,omitempty"`
	Offset       int           `json:"offset,omitempty"`
	Limit        int           `json:"limit,omitempty"`
	Observations []Observation `json:"observations,omitempty"`
}

// Observation represents data which is contained in HTTP response from FRED API.
type Observation struct {
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime"`
	Date          string `json:"date"`
	Value         string `json:"value"`
}

func (res *Res) LatestValueFloat() (float64, error) {
	latestValue := res.Observations[0].Value
	latestValueFloat64, err := strconv.ParseFloat(latestValue, 32)

	if err != nil {
		return 0, errors.New("latest value cannot be parsed at float32")
	}

	return latestValueFloat64, nil
}
