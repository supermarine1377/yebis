package request

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"supermarine1377/yebis/internal/fred/api/common"
)

const baseURL = "https://api.stlouisfed.org/fred/series/observations"

// Request represents HTTP request to get economic values from FRED API.
// Request has embedded http.Request.
type Request struct {
	httpRequest *http.Request
}

// HTTPRequest returns its embedded http.Request
func (r *Request) HTTPRequest() *http.Request {
	return r.httpRequest
}

// RequestInfo represents an optinal data to get economic values.
type RequestInfo struct {
	SeriesID      string
	RealtimeStart string
	RealtimeEnd   string
}

// NewRequest generates Request from RequestInfo.
func NewRequest(ctx context.Context, ri RequestInfo, conf common.Config) (*Request, error) {
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to url.Parse: %w", err)
	}

	v := url.Values{}
	v.Set("series_id", ri.SeriesID)
	v.Set("realtime_start", ri.RealtimeStart)
	v.Set("realtime_end", ri.RealtimeEnd)
	common.SetConfig(&v, conf)

	apiURL.RawQuery = v.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequestWithContext: %w", err)
	}

	req := &Request{httpRequest: httpReq}

	return req, nil
}
