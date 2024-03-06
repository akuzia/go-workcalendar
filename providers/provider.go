package providers

import (
	"errors"
	"strings"
	"time"
)

var ErrCodeDoesNotExist = errors.New("provided code does not exist")

type Provider struct {
	// DaysOff list of holidays and non-work days other than saturday&sunday
	DaysOff []time.Time
	// WorkDays list of working saturday&sunday
	WorkDays []time.Time
}

func GetProvider(code string) (Provider, error) {
	code = strings.ToLower(code)

	switch code {
	case "ru":
		return newRUProvider(), nil
	}

	return Provider{}, ErrCodeDoesNotExist
}
