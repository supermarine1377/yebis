// package series provides a functionality to get economic data (series) from FRED API.
package series

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"supermarine1377/yebis/internal/fred/api/common"
	common_error "supermarine1377/yebis/internal/fred/api/common/error"
	"supermarine1377/yebis/internal/fred/api/series/request"
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

// ErrFREDAPIInternalServer is an error that says FRED API returned Internal Server Error (500).
var ErrFREDAPIInternalServer = errors.New("FRED API returned Internal Server Error")

// ErrFREDAPIBadRequest is an error that says FRED API returned Bad Request Error (400).
var ErrFREDAPIBadRequest = errors.New("FRED API returned Bad Request Error")

// Get economic data via FRED API.
// The API reference: https://fred.stlouisfed.org/docs/api/fred/series.html
func Get(ctx context.Context, seriesID, realtimeStart, realtimeEnd string, config common.Config) (*Res, error) {
	req, err := request.NewRequest(
		ctx,
		request.RequestInfo{
			SeriesID:      seriesID,
			RealtimeStart: realtimeStart,
			RealtimeEnd:   realtimeEnd,
		},
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request.NewRequest: %w", err)
	}
	httpRes, err := http.DefaultClient.Do(req.HTTPRequest())
	if err != nil {
		return nil, fmt.Errorf("failed to http.Client.Do: %w", err)
	}

	defer httpRes.Body.Close()

	switch httpRes.StatusCode {
	case http.StatusBadRequest:
		er, err := common_error.ParseErrorRes(httpRes)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to error.ParseErrorRes: %w", ErrFREDAPIBadRequest, err)
		}
		return nil, fmt.Errorf("%w: %+v", ErrFREDAPIBadRequest, *er)
	case http.StatusInternalServerError:
		return nil, ErrFREDAPIInternalServer
	}

	b, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to io.ReadAll: %w", err)
	}
	var Res Res
	if err := json.Unmarshal(b, &Res); err != nil {
		return nil, fmt.Errorf("fail to json.Unmarshal: %w", err)
	}

	return &Res, nil
}
