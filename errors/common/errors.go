package common

import "errors"

var (
	ErNotFound         = errors.New("Not Found")
	ErBrokenConnection = errors.New("Broken Connection")
	ErUnknown          = errors.New("Unknown Error")
)
