package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/supermarine1377/yebis/pkg/config"
	"github.com/supermarine1377/yebis/pkg/fred/series"
	"github.com/supermarine1377/yebis/pkg/investment_score"
	"github.com/supermarine1377/yebis/pkg/investment_score/calculator"
	"github.com/supermarine1377/yebis/pkg/investment_score/record"
)

func init() {
	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)
}

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

func (a *App) CalculateInvestmentScore(ctx context.Context) (int, error) {
	dc := calculator.New(series.NewFetcher(a.config))
	c := investment_score.NewCalculator(dc)
	return c.Calculate(ctx)
}

func (a *App) Run() error {
	ctx := context.Background()
	score, err := a.CalculateInvestmentScore(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"failed to calculate investment score",
			slog.Any("error message", err),
		)
		return err
	}

	slog.InfoContext(
		ctx,
		"investment score successfully calculated",
	)
	fmt.Printf("score: %d\n", score)

	if err := record.Write(score); err != nil {
		slog.ErrorContext(
			ctx,
			"Failed to record the calculated investment score",
			slog.Any("error message", err),
		)
		return err
	}
	return nil
}
