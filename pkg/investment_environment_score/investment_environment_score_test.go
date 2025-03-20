// package investment_environment_score is responsible for calculating investment_environment_score
//
//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/$GOFILE
package investment_environment_score

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermarine1377/yebis/pkg/investment_environment_score/mock"
	"go.uber.org/mock/gomock"
)

func TestCalculator_Calculate(t *testing.T) {
	tests := []struct {
		name        string
		prepareMock func(t *testing.T, m *mock.MockInvestmentEnvironmentCalculator)
		want        int
		wantErr     bool
	}{
		{
			name: "If all the scores are 2, the score should be 10",
			prepareMock: func(t *testing.T, m *mock.MockInvestmentEnvironmentCalculator) {
				t.Helper()
				m.EXPECT().FEDFUNDS(gomock.Any()).Return(2, nil)
				m.EXPECT().T10YFF(gomock.Any()).Return(2, nil)
				m.EXPECT().US10Y(gomock.Any()).Return(2, nil)
				m.EXPECT().BAA10Y(gomock.Any()).Return(2, nil)
				m.EXPECT().USDINDEX(gomock.Any()).Return(2, nil)
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "If FEDFUNDS, T10YFF, US10Y, BAA10Y, and USDINDEX return 2, 2, -2, 2, and -2, respectively, the score should be 2",
			prepareMock: func(t *testing.T, m *mock.MockInvestmentEnvironmentCalculator) {
				t.Helper()
				m.EXPECT().FEDFUNDS(gomock.Any()).Return(2, nil)
				m.EXPECT().T10YFF(gomock.Any()).Return(2, nil)
				m.EXPECT().US10Y(gomock.Any()).Return(-2, nil)
				m.EXPECT().BAA10Y(gomock.Any()).Return(2, nil)
				m.EXPECT().USDINDEX(gomock.Any()).Return(-2, nil)
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "If FEDFUNDS returns an error, Calculate should return an error",
			prepareMock: func(t *testing.T, m *mock.MockInvestmentEnvironmentCalculator) {
				t.Helper()
				m.EXPECT().FEDFUNDS(gomock.Any()).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockInvestmentEnvironmentCalculator(ctrl)
			tt.prepareMock(t, m)
			c := NewCalculator(m)
			got, err := c.Calculate(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
