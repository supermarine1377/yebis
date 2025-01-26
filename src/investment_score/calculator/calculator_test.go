package calculator_test

import (
	"context"
	"supermarine1377/yebis/src/fred/series/response"
	"supermarine1377/yebis/src/investment_score/calculator"
	"supermarine1377/yebis/src/investment_score/calculator/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCalculator_FEDFUNDS(t *testing.T) {
	type args struct {
		ctx   context.Context
		score int
	}
	tests := []struct {
		args                     args
		name                     string
		prepareMockSeriesFetcher func(m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If today's FEDFUNDS is 0.5 and a year ago's FEDFUNDS is 0, then the score should be -2",
			args: args{
				ctx:   context.Background(),
				score: 0,
			},
			prepareMockSeriesFetcher: func(m *mock.MockSeriesFetcher) {
				// A year ago's FEDFUNDS is 0
				m.EXPECT().Fetch(gomock.Any(), "FEDFUNDS", gomock.Any()).Return(&response.Res{
					Observations: []response.Observation{
						{Value: "0"},
					},
				}, nil).After(
					// Today's FEDFUNDS is 0.5
					m.EXPECT().Fetch(gomock.Any(), "FEDFUNDS", gomock.Any()).Return(&response.Res{
						Observations: []response.Observation{
							{Value: "0.5"},
						},
					}, nil),
				)
			},
			want: -2,
		},
		{
			name: "If today's FEDFUNDS is 0 and a year ago's FEDFUNDS is 0.25, then the score should be 2",
			args: args{
				ctx:   context.Background(),
				score: 0,
			},
			prepareMockSeriesFetcher: func(m *mock.MockSeriesFetcher) {
				m.EXPECT().Fetch(gomock.Any(), "FEDFUNDS", gomock.Any()).Return(&response.Res{
					Observations: []response.Observation{
						{Value: "0.25"},
					},
				}, nil).After(
					m.EXPECT().Fetch(gomock.Any(), "FEDFUNDS", gomock.Any()).Return(&response.Res{
						Observations: []response.Observation{
							{Value: "0"},
						},
					}, nil),
				)
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockSeriesFetcher(ctrl)
			c := calculator.New(m)
			tt.prepareMockSeriesFetcher(m)

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
