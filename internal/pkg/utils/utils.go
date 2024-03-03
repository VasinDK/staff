package utils

import (
	"strconv"
)

type Utils struct{}

func New() *Utils {
	return &Utils{}
}

func (u *Utils) Atoi(a string) int {
	if a == "" {
		return 0
	}

	i, _ := strconv.Atoi(a)
	return i
}
