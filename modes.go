package mpesa

import (
	"errors"
	"strings"
)

const (
	SYNC Mode = iota
	ASYNC
)

// Mode represents the available request Mode
// requests can be either ASYNC or SYNC
type Mode int

func RequestModeStr(mode string) (Mode, error) {
	mode = strings.ToLower(mode)
	switch mode {
	case "sync", "synchronous":
		return SYNC, nil
	case "async", "asynchronous":
		return ASYNC, nil
	default:
		return 0, errors.New("invalid request mode")
	}
}

func RequestModeInt(mode int) (Mode, error) {
	switch mode {
	case 0:
		return SYNC, nil
	case 1:
		return ASYNC, nil
	default:
		return 0, errors.New("invalid request mode")
	}
}
