package pkg

import (
	"context"
	"fmt"

	"github.com/supermarine1377/yebis/pkg/config"
	"github.com/supermarine1377/yebis/pkg/fred/series"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/record"
)

type App struct {
	config *config.Config
}

func NewApp() (*App, error) {
	config, err := config.New()
	if err != nil {
		return nil, err
	}

	return &App{
		config: config,
	}, nil
}

func (a *App) CalculateInvestmentEnvironmentScore(ctx context.Context) (int, error) {
	dc := calculator.New(series.NewFetcher(a.config))
	c := investment_environment_score.NewCalculator(dc)
	return c.Calculate(ctx)
}

func (a *App) Run() error {
	ctx := context.Background()
	score, err := a.CalculateInvestmentEnvironmentScore(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("score: %d\n", score)

	if err := record.Write(score); err != nil {
		return err
	}
	return nil
}
