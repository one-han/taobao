package api

import (
	"net/url"
)

type TaobaoRequest interface {
	// Check() error
	// GetTextParams() (string, error)
	GetApiMethodName() string
	SetValue(key, value string)
	GetValues() url.Values
}
