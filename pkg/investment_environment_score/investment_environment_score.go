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

	fredfunds, err := c.iec.FEDFUNDS(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of FEDFUNDS: %w", err)
	}
	fmt.Printf("fredfunds: %d\n", fredfunds)
	score += fredfunds

	t10yff, err := c.iec.T10YFF(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of T10YFF: %w", err)
	}
	fmt.Println("t10yff: ", t10yff)
	score += t10yff

	us10y, err := c.iec.US10Y(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of US10Y: %w", err)
	}
	fmt.Println("us10y: ", us10y)
	score += us10y

	baa10y, err := c.iec.BAA10Y(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of BAA10Y: %w", err)
	}
	fmt.Println("baa10y: ", baa10y)
	score += baa10y

	usdindex, err := c.iec.USDINDEX(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate a part of USDINDEX: %w", err)
	}
	fmt.Println("usdindex: ", usdindex)
	score += usdindex

	return score, nil
}
