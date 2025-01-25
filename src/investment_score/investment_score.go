// package investment_score is responsble for calculating investment_score
//
//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/$GOFILE
package investment_score

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

type Calculator struct {
	pc partCalculator
}

// diffCalculator calculates a diffence of each ecnomic data between that of today and that of a year ago
type partCalculator interface {
	FEDFUNDS(ctx context.Context) (float64, error)
	US10Y(ctx context.Context) (float64, error)
	T10YFF(ctx context.Context) (float64, error)
	BAA10Y(ctx context.Context) (float64, error)
	USDINDEX(ctx context.Context) (float64, error)
}

func NewCalculator(pc partCalculator) *Calculator {
	return &Calculator{
		pc: pc,
	}
}

func (c *Calculator) Do(ctx context.Context) (int, error) {
	var score int

	slog.InfoContext(ctx, "calculating a part of FEDFUNDS...")
	fedfundsDiff, err := c.pc.FEDFUNDS(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of FEDFUNDS: %w", err)
	}

	slog.InfoContext(ctx, "successfully calculated a part of FEDFUNDS", slog.Float64("value", fedfundsDiff))
	if fedfundsDiff > 0.25 {
		score = score - 2
	} else {
		score = score + 2
	}

	slog.Info("completed 1/5 part of calculateing investment score", slog.Int("score at this point", score))

	slog.InfoContext(ctx, "calculating a part of T10YFF...")
	t10yff, err := c.pc.T10YFF(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of T10YFF: %w", err)
	}
	slog.InfoContext(ctx, "successfully calculated a part of T10YFF", slog.Float64("value", t10yff))
	if t10yff >= 1 {
		score = score + 2
	} else if t10yff < 0 {
		score = score - 2
	}

	slog.Info("completed 2/5 part of calculateing investment score", slog.Int("score at this point", score))

	slog.InfoContext(ctx, "calculating a part of US10Y...")
	us10yDiff, err := c.pc.US10Y(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of US10Y: %w", err)
	}
	slog.InfoContext(ctx, "successfully calculated a part of US10Y", slog.Float64("value", us10yDiff))
	if us10yDiff < 0 {
		score = score - 2
	} else {
		score = score + 2
	}

	slog.Info("completed 3/5 part of calculateing investment score", slog.Int("score at this point", score))

	slog.InfoContext(ctx, "calculating a part of BAA10Y...")
	baa10yDiff, err := c.pc.BAA10Y(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of BAA10Y: %w", err)
	}
	slog.InfoContext(ctx, "successfully calculated a part of BAA10Y", slog.Float64("value", baa10yDiff))
	if baa10yDiff > 0 {
		score = score - 2
	} else {
		score = score + 2
	}

	slog.Info("completed 4/5 part of calculateing investment score", slog.Int("score at this point", score))

	slog.InfoContext(ctx, "calculating a part of USDINDEX...")
	usdDiff, err := c.pc.USDINDEX(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of USDINDEX: %w", err)
	}
	slog.InfoContext(ctx, "successfully calculated a part of USDINDEX", slog.Float64("value", usdDiff))
	if usdDiff > 0 {
		score = score - 2
	} else {
		score = score + 2
	}

	slog.Info("completed 5/5 part of calculateing investment score", slog.Int("score at this point", score))

	return score, nil
}
