package investment_score_test

import (
	"context"
	"supermarine1377/yebis/src/investment_score"
	mock_investment_score "supermarine1377/yebis/src/investment_score/mock/investment_score"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCalculator_Do(t *testing.T) {
	tests := []struct {
		name                          string
		want                          int
		prepareMockMockdiffCalculator func(ctx context.Context, m *mock_investment_score.MockpartCalculator)
		wantErr                       bool
	}{
		{
			name: "1st",
			want: 2,
			prepareMockMockdiffCalculator: func(ctx context.Context, m *mock_investment_score.MockpartCalculator) {
				m.EXPECT().FEDFUNDS(ctx).Return(float64(0.13), nil)
				m.EXPECT().T10YFF(ctx).Return(float64(1.93), nil)
				m.EXPECT().US10Y(ctx).Return(float64(0.5), nil)
				m.EXPECT().BAA10Y(ctx).Return(float64(0.04), nil)
				m.EXPECT().USDINDEX(ctx).Return(float64(1), nil)
			},
		},
		{
			name: "2nd",
			want: -6,
			prepareMockMockdiffCalculator: func(ctx context.Context, m *mock_investment_score.MockpartCalculator) {
				m.EXPECT().FEDFUNDS(ctx).Return(float64(0.3), nil)
				m.EXPECT().T10YFF(ctx).Return(float64(-1), nil)
				m.EXPECT().US10Y(ctx).Return(float64(-1), nil)
				m.EXPECT().BAA10Y(ctx).Return(float64(-1), nil)
				m.EXPECT().USDINDEX(ctx).Return(float64(1), nil)
			},
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		t.Run(tt.name, func(t *testing.T) {
			m := mock_investment_score.NewMockpartCalculator(ctrl)

			c := investment_score.NewCalculator(m)
			ctx := context.Background()

			tt.prepareMockMockdiffCalculator(ctx, m)

			got, err := c.Do(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
