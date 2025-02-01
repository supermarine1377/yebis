package request

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/supermarine1377/yebis/pkg/fred/common"

	"golang.org/x/exp/slog"
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
type Option struct {
	SeriesID       string
	ObservationEnd string
}

// NewRequest generates Request from RequestInfo.
func NewRequest(ctx context.Context, o *Option, conf common.Config) (*Request, error) {
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to url.Parse: %w", err)
	}

	v := url.Values{}
	v.Set("series_id", o.SeriesID)
	v.Set("observation_end", o.ObservationEnd)
	v.Set("limit", "1")
	v.Set("order_by", "observation_date")
	v.Set("sort_order", "desc")

	common.SetConfig(&v, conf)

	apiURL.RawQuery = v.Encode()

	slog.Info(
		"HTTP request to FEDFUNDS",
		slog.String("URL", apiURL.String()),
	)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequestWithContext: %w", err)
	}

	req := &Request{httpRequest: httpReq}

	return req, nil
}
