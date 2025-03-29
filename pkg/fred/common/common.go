//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/$GOFILE
package common

import (
	"net/url"
)

// Config represents configuration for this application.
type Config interface {
	FEDAPIKEY() string
}

// SetConfig is a function that configures the URL of the HTTP request.
// It sets the 'api_key' and 'file_type' fields, where 'api_key' is obtained from the given configuration
// and 'file_type' is set to 'json'.
func SetConfig(url *url.Values, config Config) {
	url.Set("api_key", config.FEDAPIKEY())
	url.Set("file_type", "json")
}
