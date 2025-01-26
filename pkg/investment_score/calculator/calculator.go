//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/$GOFILE
package calculator

import (
	"context"
	"errors"
	"fmt"
	"supermarine1377/yebis/pkg/fred/series/response"
	"supermarine1377/yebis/pkg/fred/series/series_id"
	"time"
)

type Calculator struct {
	sf SeriesFetcher
}
type SeriesFetcher interface {
	Fetch(ctx context.Context, seriesID string, obeservationEnd time.Time) (*response.Res, error)
}

func New(sf SeriesFetcher) *Calculator {
	return &Calculator{sf: sf}
}

// FEDFUNDS calculates the score of FEDFUNDS.
// If the difference between today's FEDFUNDS and that of a year ago is greater than 0.25, the score is decreased by 2.
// Otherwise, the score is increased by 2.
func (c *Calculator) FEDFUNDS(ctx context.Context, score int) (int, error) {
	diff, err := c.diff(ctx, series_id.FEDFUNDS)
	if err != nil {
		return 0, err
	}
	if diff > 0.25 {
		score = score - 2
	} else {
		score = score + 2
	}
	return score, nil
}

func (c *Calculator) US10Y(ctx context.Context, score int) (int, error) {
	diff, err := c.diff(ctx, series_id.US10Y)
	if err != nil {
		return 0, err
	}
	if diff < 0 {
		score = score - 2
	} else {
		score = score + 2
	}
	return score, nil
}

func (c *Calculator) T10YFF(ctx context.Context, score int) (int, error) {
	res, err := c.sf.Fetch(
		ctx,
		series_id.T10YFF,
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	val, err := res.LatestValueFloat()
	if err != nil {
		return 0, err
	}
	if val >= 1 {
		score = score + 2
	} else if val < 0 {
		score = score - 2
	}
	return score, nil
}

func (c *Calculator) BAA10Y(ctx context.Context, score int) (int, error) {
	diff, err := c.diff(ctx, series_id.BAA10Y)
	if err != nil {
		return 0, err
	}
	if diff > 0 {
		score = score - 2
	} else {
		score = score + 2
	}
	return score, nil
}

func (c *Calculator) USDINDEX(ctx context.Context, score int) (int, error) {
	diff, err := c.diff(ctx, series_id.USDINDEX)
	if err != nil {
		return 0, err
	}
	if diff > 0 {
		score = score - 2
	} else {
		score = score + 2
	}

	return score, nil
}

func (c *Calculator) diff(ctx context.Context, seriesID string) (float64, error) {
	now := time.Now()
	today, err := c.sf.Fetch(
		ctx,
		seriesID,
		now,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s of today: %w", seriesID, err)
	}

	yearago, err := c.sf.Fetch(
		ctx,
		seriesID,
		now.AddDate(-1, 0, 0),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s of a year ago: %w", seriesID, err)
	}
	diff, err := diff(today, yearago)
	if err != nil {
		return 0, fmt.Errorf("failed to caliculate diff between %s of today and that of a year ago: %w", seriesID, err)
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
