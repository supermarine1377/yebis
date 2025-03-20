// package series provides a functionality to get economic data (series) from FRED API.
package series

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/supermarine1377/yebis/pkg/fred/common"
	common_error "github.com/supermarine1377/yebis/pkg/fred/common/error"
	"github.com/supermarine1377/yebis/pkg/fred/series/request"
	"github.com/supermarine1377/yebis/pkg/fred/series/response"
)

// ErrFREDAPIInternalServer is an error that says FRED API returned Internal Server Error (500).
var ErrFREDAPIInternalServer = errors.New("FRED API returned Internal Server Error")

// ErrFREDAPIBadRequest is an error that says FRED API returned Bad Request Error (400).
var ErrFREDAPIBadRequest = errors.New("FRED API returned Bad Request Error")

type Fetcher struct {
	config    common.Config
	transport http.RoundTripper
}

func NewFetcher(config common.Config) *Fetcher {
	return &Fetcher{
		config:    config,
		transport: http.DefaultTransport,
	}
}

func (f *Fetcher) SetTransport(transport http.RoundTripper) {
	f.transport = transport
}

// Fetch economic data via FRED API.
// The API reference: https://fred.stlouisfed.org/docs/api/fred/series.html
func (f *Fetcher) Fetch(ctx context.Context, seriesID string, obeservationEnd time.Time) (*response.Res, error) {
	var (
		res *response.Res
		err error
	)
	for {
		res, err = f.get(ctx, seriesID, obeservationEnd)
		if err != nil {
			return nil, err
		}
		if _, err := res.LatestValueFloat(); err != nil {
			obeservationEnd = obeservationEnd.AddDate(0, 0, -1)
			continue
		}
		break
	}
	return res, err
}

func (f *Fetcher) get(ctx context.Context, seriesID string, obeservationEnd time.Time) (*response.Res, error) {
	req, err := request.NewRequest(
		ctx,
		&request.Option{
			SeriesID:       seriesID,
			ObservationEnd: Date(obeservationEnd),
		},
		f.config,
	)
	if err != nil {
		return nil, err
	}
	httpRes, err := f.transport.RoundTrip(req.HTTPRequest())
	if err != nil {
		return nil, err
	}

	defer httpRes.Body.Close()

	switch httpRes.StatusCode {
	case http.StatusBadRequest:
		er, err := common_error.ParseErrorRes(httpRes)
		if err == nil {
			return nil, fmt.Errorf("%w: message:%s, error_code: %d", ErrFREDAPIBadRequest, er.ErrorMessage, er.ErrorCode)
		}
		return nil, ErrFREDAPIBadRequest
	case http.StatusInternalServerError:
		return nil, ErrFREDAPIInternalServer
	}

	b, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	var Res response.Res
	if err := json.Unmarshal(b, &Res); err != nil {
		return nil, err
	}

	return &Res, nil
}
