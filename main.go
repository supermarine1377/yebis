package main

import (
	"context"
	"fmt"
	"os"
	"supermarine1377/yebis/internal/investment_score"
	"supermarine1377/yebis/internal/investment_score/economic_data"
	"supermarine1377/yebis/internal/investment_score/record"
	"supermarine1377/yebis/internal/pkg/config"

	"golang.org/x/exp/slog"
)

func init() {
	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)
}

const panicMessage = "An unexpected error has occurred."

func main() {
	slog.Info("start to calculate Investment Score...")

	ctx := context.Background()
	config, err := config.NewConfig()
	if err != nil {
		slog.ErrorContext(
			ctx,
			"failed to load configration",
			slog.Any("error message", err),
		)
		panic(panicMessage)
	}

	dc := economic_data.NewCalculator(config)
	c := investment_score.NewCalculator(dc)
	score, err := c.Do(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"failed to calculate investment score",
			slog.Any("error message", err),
		)
		panic(panicMessage)
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
		panic(err)
	}
}
