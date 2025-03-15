// package investment_environment_score is responsible for calculating investment_environment_score
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
	FEDFUNDS(ctx context.Context) (int, error)
	US10Y(ctx context.Context) (int, error)
	T10YFF(ctx context.Context) (int, error)
	BAA10Y(ctx context.Context) (int, error)
	USDINDEX(ctx context.Context) (int, error)
}

func NewCalculator(iec InvestmentEnvironmentCalculator) *Calculator {
	return &Calculator{
		iec: iec,
	}
}

func (c *Calculator) Calculate(ctx context.Context) (int, error) {
	var score int

	parts := []struct {
		name string
		fn   func(ctx context.Context) (int, error)
	}{
		{"FEDFUNDS", c.iec.FEDFUNDS},
		{"T10YFF", c.iec.T10YFF},
		{"US10Y", c.iec.US10Y},
		{"BAA10Y", c.iec.BAA10Y},
		{"USDINDEX", c.iec.USDINDEX},
	}

	for _, part := range parts {
		value, err := part.fn(ctx)
		if err != nil {
			return 0, fmt.Errorf("failed to calculate a part of %s: %w", part.name, err)
		}
		fmt.Printf("%s: %d\n", part.name, value)
		score += value
	}

	return score, nil
}
