// package investment_environment_score is responsble for calculating investment_environment_score
//
//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/$GOFILE
package investment_environment_score

import (
	"context"
	"fmt"
)

type Calculator struct {
	iec InvestmentEnvironmentCalculator
}

type InvestmentEnvironmentCalculator interface {
	FEDFUNDS(ctx context.Context, score int) (int, error)
	US10Y(ctx context.Context, score int) (int, error)
	T10YFF(ctx context.Context, score int) (int, error)
	BAA10Y(ctx context.Context, score int) (int, error)
	USDINDEX(ctx context.Context, score int) (int, error)
}

func NewCalculator(iec InvestmentEnvironmentCalculator) *Calculator {
	return &Calculator{
		iec: iec,
	}
}

func (c *Calculator) Calculate(ctx context.Context) (int, error) {
	var score int

	score, err := c.iec.FEDFUNDS(ctx, score)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of FEDFUNDS: %w", err)
	}

	score, err = c.iec.T10YFF(ctx, score)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of T10YFF: %w", err)
	}

	score, err = c.iec.US10Y(ctx, score)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of US10Y: %w", err)
	}

	score, err = c.iec.BAA10Y(ctx, score)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of BAA10Y: %w", err)
	}

	score, err = c.iec.USDINDEX(ctx, score)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of USDINDEX: %w", err)
	}

	return score, nil
}
