package utils

import (
	"errors"
)

var (
	localIPPrefix = [...]string{"192.168", "10.0", "169.254", "172.16"}
	ErrPort       = errors.New("invalid port")
)
