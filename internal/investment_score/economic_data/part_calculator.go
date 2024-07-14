package economic_data

import (
	"context"
	"errors"
	"fmt"
	"supermarine1377/yebis/internal/fred/api/common"
	"supermarine1377/yebis/internal/fred/api/series"
	"supermarine1377/yebis/internal/fred/api/series/response"
	"supermarine1377/yebis/internal/fred/api/series/series_id"
	"time"
)

type Calculator struct {
	config common.Config
}

func NewCalculator(config common.Config) *Calculator {
	return &Calculator{config: config}
}

func (c *Calculator) FEDFUNDS(ctx context.Context) (float64, error) {
	return c.diff(ctx, series_id.FEDFUNDS)
}

func (c *Calculator) US10Y(ctx context.Context) (float64, error) {
	return c.diff(ctx, series_id.US10Y)
}

func (c *Calculator) T10YFF(ctx context.Context) (float64, error) {
	res, err := series.Get(
		ctx,
		series_id.T10YFF,
		time.Now(),
		c.config,
	)
	if err != nil {
		return 0, err
	}
	val, err := res.LatestValueFloat()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *Calculator) BAA10Y(ctx context.Context) (float64, error) {
	return c.diff(ctx, series_id.BAA10Y)
}

func (c *Calculator) USDINDEX(ctx context.Context) (float64, error) {
	return c.diff(ctx, series_id.USDINDEX)
}

func (c *Calculator) diff(ctx context.Context, dataName string) (float64, error) {
	now := time.Now()
	today, err := series.Get(
		ctx,
		dataName,
		now,
		c.config,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s of today: %w", dataName, err)
	}

	yearago, err := series.Get(
		ctx,
		dataName,
		now.AddDate(-1, 0, 0),
		c.config,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s of a year ago: %w", dataName, err)
	}
	diff, err := diff(today, yearago)
	if err != nil {
		return 0, fmt.Errorf("failed to caliculate diff between %s of today and that of a year ago: %w", dataName, err)
	}
	return diff, nil
}

var errNumberOfObservationsMustBeOne = errors.New("number of observations must one")

// If data of today or that of one year ago contains more than one Observersion, this function threw error.
// Then substract return a defference between obeservation value of res1 and that of res2.
func diff(res1 *response.Res, res2 *response.Res) (float64, error) {
	if len(res1.Observations) != 1 || len(res2.Observations) != 1 {
		return 0, errNumberOfObservationsMustBeOne
	}
	val1, err := res1.LatestValueFloat()
	if err != nil {
		return 0, err
	}
	val2, err := res2.LatestValueFloat()
	if err != nil {
		return 0, err
	}
	return val1 - val2, nil
}
