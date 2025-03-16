package calculator_test

import (
	"context"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermarine1377/yebis/pkg/fred/series/response"
	"github.com/supermarine1377/yebis/pkg/fred/series/series_id"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/calculator/mock"
	"go.uber.org/mock/gomock"
)

func configureMockSeriesFetcherWithPastAndPresent(t *testing.T, m *mock.MockSeriesFetcher, seriesID string, ayearago, today string) {
	t.Helper()

	m.EXPECT().Fetch(gomock.Any(), seriesID, gomock.Any()).Return(&response.Res{
		Observations: []response.Observation{
			{Value: ayearago},
		},
	}, nil).After(
		m.EXPECT().Fetch(gomock.Any(), seriesID, gomock.Any()).Return(&response.Res{
			Observations: []response.Observation{
				{Value: today},
			},
		}, nil),
	)
}

func TestCalculator_FEDFUNDS(t *testing.T) {
	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If today's FEDFUNDS is 0.5 and a year ago's FEDFUNDS is 0, then the score should be -2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.FEDFUNDS, "0", "0.5")
			},
			want: -2,
		},
		{
			name: "If today's FEDFUNDS is 0 and a year ago's FEDFUNDS is 0.25, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.FEDFUNDS, "0.25", "0")
			},
			want: 2,
		},
		{
			name: "異常系",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				m.EXPECT().Fetch(gomock.Any(), series_id.FEDFUNDS, gomock.Any()).Return(nil, assert.AnError)
			},
			wantErr: true,
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
	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If today's US10Y is 1 and a year ago's US10Y is 1.5, then the score should be -2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.US10Y, "1.5", "1")
			},
			want:    -2,
			wantErr: false,
		},
		{
			name: "If today's US10Y is 1.5 and a year ago's US10Y is 1, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.US10Y, "1", "1.5")
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

func TestCalculator_T10YFF(t *testing.T) {
	prepateMockSeriesFetcher := func(t *testing.T, m *mock.MockSeriesFetcher, value string) {
		t.Helper()
		m.EXPECT().Fetch(gomock.Any(), series_id.T10YFF, gomock.Any()).Return(&response.Res{
			Observations: []response.Observation{
				{Value: value},
			},
		}, nil)
	}

	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If T10YFF is 1, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepateMockSeriesFetcher(t, m, "1")
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "If T10YFF is -1, then the score should be -2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepateMockSeriesFetcher(t, m, "-1")
			},
			want:    -2,
			wantErr: false,
		},
		{
			name: "If T10YFF is 0.5, then the score should be 0",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				prepateMockSeriesFetcher(t, m, "0.5")
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "異常系",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				m.EXPECT().Fetch(gomock.Any(), series_id.T10YFF, gomock.Any()).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "異常系 - float64の最大値",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				val := strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64)
				prepateMockSeriesFetcher(t, m, val)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockSeriesFetcher(ctrl)
			tt.prepareMockSeriesFetcher(t, m)
			c := calculator.New(m)

			got, err := c.T10YFF(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculator_BAA10Y(t *testing.T) {
	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If BAA10Y is 1 and a year ago's BAA10Y is -1, then the score should be -2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.BAA10Y, "-1", "1")
			},
			want: -2,
		},
		{
			name: "If BAA10Y is -1 and a year ago's BAA10Y is 1, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.BAA10Y, "1", "-1")
			},
			want: 2,
		},
		{
			name: "異常系",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				m.EXPECT().Fetch(gomock.Any(), series_id.BAA10Y, gomock.Any()).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockSeriesFetcher(ctrl)
			tt.prepareMockSeriesFetcher(t, m)
			c := calculator.New(m)

			got, err := c.BAA10Y(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculator_USDINDEX(t *testing.T) {
	tests := []struct {
		name                     string
		prepareMockSeriesFetcher func(t *testing.T, m *mock.MockSeriesFetcher)
		want                     int
		wantErr                  bool
	}{
		{
			name: "If USDINDEX is 1 and a year ago's USDINDEX is -1, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.USDINDEX, "-1", "1")
			},
			want: -2,
		},
		{
			name: "If USDINDEX is -1 and a year ago's USDINDEX is 1, then the score should be 2",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				configureMockSeriesFetcherWithPastAndPresent(t, m, series_id.USDINDEX, "1", "-1")
			},
			want: 2,
		},
		{
			name: "異常系",
			prepareMockSeriesFetcher: func(t *testing.T, m *mock.MockSeriesFetcher) {
				m.EXPECT().Fetch(gomock.Any(), series_id.USDINDEX, gomock.Any()).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockSeriesFetcher(ctrl)
			tt.prepareMockSeriesFetcher(t, m)
			c := calculator.New(m)

			got, err := c.USDINDEX(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
