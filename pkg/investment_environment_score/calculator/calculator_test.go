package calculator_test

import (
	"context"
	"testing"

	"github.com/supermarine1377/yebis/pkg/fred/series/response"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func prepareMockSeriesFetcher(t *testing.T, m *mock.MockSeriesFetcher, ayearago, today string) {
	t.Helper()

	m.EXPECT().Fetch(gomock.Any(), "FEDFUNDS", gomock.Any()).Return(&response.Res{
		Observations: []response.Observation{
			{Value: ayearago},
		},
	}, nil).After(
		m.EXPECT().Fetch(gomock.Any(), "FEDFUNDS", gomock.Any()).Return(&response.Res{
			Observations: []response.Observation{
				{Value: today},
			},
		}, nil),
	)
}
func TestCalculator_FEDFUNDS(t *testing.T) {

	type args struct {
		ctx   context.Context
		score int
	}
	tests := []struct {
		args                     args
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If today's FEDFUNDS is 0.5 and a year ago's FEDFUNDS is 0, then the score should be -2",
			args: args{
				ctx:   context.Background(),
				score: 0,
			},
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepareMockSeriesFetcher(t, m, "0", "0.5")
			},
			want: -2,
		},
		{
			name: "If today's FEDFUNDS is 0 and a year ago's FEDFUNDS is 0.25, then the score should be 2",
			args: args{
				ctx:   context.Background(),
				score: 0,
			},
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepareMockSeriesFetcher(t, m, "0.25", "0")
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockSeriesFetcher(ctrl)
			c := calculator.New(m)
			tt.prepareMockSeriesFetcher(t, m)

			got, err := c.FEDFUNDS(tt.args.ctx, tt.args.score)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
