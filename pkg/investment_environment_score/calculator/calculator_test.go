package calculator_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermarine1377/yebis/pkg/fred/series/response"
	"github.com/supermarine1377/yebis/pkg/fred/series/series_id"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator/mock"
	"go.uber.org/mock/gomock"
)

func TestCalculator_FEDFUNDS(t *testing.T) {
	prepareMockSeriesFetcher := func(t *testing.T, m *mock.MockSeriesFetcher, ayearago, today string) {
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
	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If today's FEDFUNDS is 0.5 and a year ago's FEDFUNDS is 0, then the score should be -2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepareMockSeriesFetcher(t, m, "0", "0.5")
			},
			want: -2,
		},
		{
			name: "If today's FEDFUNDS is 0 and a year ago's FEDFUNDS is 0.25, then the score should be 2",
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

			got, err := c.FEDFUNDS(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculator_US10Y(t *testing.T) {
	prepareMockSeriesFetcher := func(t *testing.T, m *mock.MockSeriesFetcher, ayearago, today string) {
		t.Helper()
		m.EXPECT().Fetch(gomock.Any(), series_id.US10Y, gomock.Any()).Return(&response.Res{
			Observations: []response.Observation{
				{Value: ayearago},
			},
		}, nil).After(
			m.EXPECT().Fetch(gomock.Any(), series_id.US10Y, gomock.Any()).Return(&response.Res{
				Observations: []response.Observation{
					{Value: today},
				},
			}, nil),
		)
	}

	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If today's US10Y is 1 and a year ago's US10Y is 1.5, then the score should be -2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepareMockSeriesFetcher(t, m, "1.5", "1")
			},
			want:    -2,
			wantErr: false,
		},
		{
			name: "If today's US10Y is 1.5 and a year ago's US10Y is 1, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepareMockSeriesFetcher(t, m, "1", "1.5")
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "異常系",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				m.EXPECT().Fetch(gomock.Any(), series_id.US10Y, gomock.Any()).Return(nil, assert.AnError)
			},
			wantErr: true,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockSeriesFetcher(ctrl)
			c := calculator.New(m)
			tt.prepareMockSeriesFetcher(t, m)

			got, err := c.US10Y(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
