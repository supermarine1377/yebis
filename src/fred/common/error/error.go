package error

import (
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("failed to io.ReadAll: %w", err)
	}

	var Res ErrorRes
	if err := json.Unmarshal(b, &Res); err != nil {
		return nil, fmt.Errorf("fail to json.Unmarshal: %w", err)
	}

	return &Res, nil
}
