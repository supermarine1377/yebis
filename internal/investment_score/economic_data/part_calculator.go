package economic_data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"supermarine1377/yebis/internal/fred/api/common"
	"supermarine1377/yebis/internal/fred/api/series"
	"supermarine1377/yebis/internal/fred/api/series/frequency"
	"supermarine1377/yebis/internal/fred/api/series/series_id"
)

type DiffCalculator struct {
	config common.Config
}

func NewDiffCalculator(config common.Config) *DiffCalculator {
	return &DiffCalculator{config: config}
}

func (dc *DiffCalculator) FEDFUNDS(ctx context.Context) (float64, error) {
	return dc.do(ctx, series_id.FEDFUNDS, frequency.Monthly)
}

func (dc *DiffCalculator) US10Y(ctx context.Context) (float64, error) {
	return dc.do(ctx, series_id.US10Y, frequency.Monthly)
}

// At the edge, T10YFF becomes a discontinuous function.
// For example, on New Year's Day, T10YFF recorded by FEDFUNDS becomes somewhat discontinuous..
// So here we calculate T10YFF using US10Y and FEDFUNDS RATE
func (dc *DiffCalculator) T10YFF(ctx context.Context) (float64, error) {
	t10yff, err := dc.do(
		ctx,
		series_id.T10YFF,
		frequency.Monthly,
	)
	if err != nil {
		return 0, err
	}
	return t10yff, nil
}

func (dc *DiffCalculator) BAA10Y(ctx context.Context) (float64, error) {
	return dc.do(ctx, series_id.BAA10Y, frequency.Monthly)
}

func (dc *DiffCalculator) USDINDEX(ctx context.Context) (float64, error) {
	return dc.do(ctx, series_id.USDINDEX, frequency.Weekly)
}

func (dc *DiffCalculator) do(ctx context.Context, dataName string, freq string) (float64, error) {
	today, err := series.Get(
		ctx,
		dataName,
		series.DateToday(),
		freq,
		dc.config,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s of today: %w", dataName, err)
	}

	yearago, err := series.Get(
		ctx,
		dataName,
		series.DateYearAgo(),
		freq,
		dc.config,
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
func diff(res1 *series.Res, res2 *series.Res) (float64, error) {
	if len(res1.Observations) != 1 || len(res2.Observations) != 1 {
		return 0, errNumberOfObservationsMustBeOne
	}
	val1, err := strconv.ParseFloat(res1.Observations[0].Value, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to strconv.ParseFloat: %w", err)
	}
	val2, err := strconv.ParseFloat(res2.Observations[0].Value, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to strconv.ParseFloat: %w", err)
	}
	return val1 - val2, nil
}
