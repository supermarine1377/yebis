package src

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"supermarine1377/yebis/src/config"
	"supermarine1377/yebis/src/investment_score"
	"supermarine1377/yebis/src/investment_score/economic_data"
	"supermarine1377/yebis/src/investment_score/record"
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

func (a *App) Run() error {
	ctx := context.Background()
	dc := economic_data.NewCalculator(a.config)
	c := investment_score.NewCalculator(dc)
	score, err := c.Do(ctx)
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
			err,
		)
		return err
	}
	return nil
}
