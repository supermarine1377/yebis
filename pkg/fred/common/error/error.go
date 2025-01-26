package error

import (
	"encoding/json"
	"io"
	"net/http"
)

// ErrorRes represents an error HTTP response from FRED API.
type ErrorRes struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func ParseErrorRes(httpRes *http.Response) (*ErrorRes, error) {
	b, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	var Res ErrorRes
	if err := json.Unmarshal(b, &Res); err != nil {
		return nil, err
	}
	return &Res, nil
}
