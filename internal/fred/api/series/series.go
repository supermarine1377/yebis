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
	"supermarine1377/yebis/internal/fred/api/series/response"
	"time"
)

// ErrFREDAPIInternalServer is an error that says FRED API returned Internal Server Error (500).
var ErrFREDAPIInternalServer = errors.New("FRED API returned Internal Server Error")

// ErrFREDAPIBadRequest is an error that says FRED API returned Bad Request Error (400).
var ErrFREDAPIBadRequest = errors.New("FRED API returned Bad Request Error")

// Get economic data via FRED API.
// The API reference: https://fred.stlouisfed.org/docs/api/fred/series.html
func Get(ctx context.Context, seriesID string, obeservationEnd time.Time, config common.Config) (*response.Res, error) {
	var (
		res *response.Res
		err error
	)
	for {
		fmt.Println(obeservationEnd)
		res, err = get(ctx, seriesID, obeservationEnd, config)
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

func get(ctx context.Context, seriesID string, obeservationEnd time.Time, config common.Config) (*response.Res, error) {
	req, err := request.NewRequest(
		ctx,
		&request.Option{
			SeriesID:       seriesID,
			ObservationEnd: Date(obeservationEnd),
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
	var Res response.Res
	if err := json.Unmarshal(b, &Res); err != nil {
		return nil, fmt.Errorf("fail to json.Unmarshal: %w", err)
	}

	return &Res, nil
}
