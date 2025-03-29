// package series provides a functionality to get economic data (series) from FRED API.
package series

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/yebis/pkg/fred/common"
	"github.com/supermarine1377/yebis/pkg/fred/common/mock"
	"github.com/supermarine1377/yebis/pkg/fred/series/response"
	"go.uber.org/mock/gomock"
)

type MockRoundTripper struct {
	res *http.Response
	err error
}

func (m *MockRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return m.res, m.err
}

func mockConfig(t *testing.T) common.Config {
	t.Helper()
	ctrl := gomock.NewController(t)
	mc := mock.NewMockConfig(ctrl)
	mc.EXPECT().FEDAPIKEY().Return("test")
	return mc
}

func TestFetcher_get(t *testing.T) {
	const res = `{"units":"","output_type":1,"file_type":"json","order_by":"","sort_by":"","count":1,"offset":0,"limit":1000,"observations":[{"realtime_start":"2023-01-01","realtime":"2023-01-01","date":"2023-01-01","value":"100"}]}`

	type fields struct {
		config    common.Config
		transport http.RoundTripper
	}
	tests := []struct {
		name    string
		fields  fields
		want    *response.Res
		wantErr bool
	}{
		{
			name: "RoundTrip returns non-nil error",
			fields: fields{
				config: mockConfig(t),
				transport: &MockRoundTripper{
					res: nil,
					err: assert.AnError,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "RoundTrip returns 400 response",
			fields: fields{
				config: mockConfig(t),
				transport: &MockRoundTripper{
					res: &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(&bytes.Buffer{}),
					},
					err: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "RoundTrip returns 500 response",
			fields: fields{
				config: mockConfig(t),
				transport: &MockRoundTripper{
					res: &http.Response{
						StatusCode: http.StatusInternalServerError,
						Body:       io.NopCloser(&bytes.Buffer{}),
					},
					err: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "RoundTrip returns 200 response",
			fields: fields{
				config: mockConfig(t),
				transport: &MockRoundTripper{
					res: &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(res)),
					},
					err: nil,
				},
			},
			want: &response.Res{
				OutputType: 1,
				FileType:   "json",
				Count:      1,
				Limit:      1000,
				Observations: []response.Observation{
					{
						RealtimeStart: "2023-01-01",
						RealtimeEnd:   "2023-01-01",
						Date:          "2023-01-01",
						Value:         "100",
					},
				},
			},
			wantErr: false,
		},
	}
	const seriesID = "test"
	obeservationEnd := time.Time{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFetcher(tt.fields.config)
			f.SetTransport(tt.fields.transport)
			got, err := f.get(context.Background(), seriesID, obeservationEnd)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
