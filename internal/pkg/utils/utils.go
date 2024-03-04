// Пакет с утилитами для работы
package utils

import (
	"strconv"
)

type Utils struct{}

func New() *Utils {
	return &Utils{}
}

// Делает из string -> int с отсечением проверки на ошибки
func (u *Utils) Atoi(a string) int {
	if a == "" {
		return 0
	}

	i, _ := strconv.Atoi(a)
	return i
}
