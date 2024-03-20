// Staff пакет обработчиков запросов
package staff

import (
	"log/slog"
	"this_module/internal/usecase"
)

type Staff struct {
	uc *usecase.StaffUC
	l  *slog.Logger
}

func New(l *slog.Logger, uc *usecase.StaffUC) *Staff {
	return &Staff{
		uc,
		l,
	}
}
