package main

import (
	"context"
	"os"
	"supermarine1377/yebis/internal/fred/api/series"
	"supermarine1377/yebis/internal/fred/api/series/series_id"
	"supermarine1377/yebis/internal/pkg/config"

	"golang.org/x/exp/slog"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}

const panicMessage = "An unexpected error has occurred."

func main() {
	slog.Info("start to calculate Investment Score...")

	ctx := context.Background()
	config, err := config.NewConfig()
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		panic(panicMessage)
	}

	g, err := series.Get(ctx, series_id.FEDFUNDS, "2023-01-01", "2023-12-31", config)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		panic(panicMessage)
	}
	slog.Info("result", g)
}
