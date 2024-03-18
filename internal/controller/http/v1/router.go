// v1 HTTP роутер работающий на фреймворке chi
package v1

import (
	"log/slog"
	"this_module/internal/controller/http/v1/staff"
	"this_module/internal/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(l *slog.Logger, uc *usecase.StaffUC) (*chi.Mux, error) {
	r := chi.NewRouter()
	staff := staff.New(l, uc)

	// Восстановление после сбоя
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		// Получает сотрудников
		r.Get("/staff/{companyId}", staff.Get)

		// Добавляет сотрудника
		r.Post("/staff", staff.Add)

		// Удаляет сотрудника по id
		r.Delete("/staff/{deleteId}", staff.Delete)

		// Обновляет поля сотрудника по id
		r.Put("/staff", staff.Update)
	})

	return r, nil
}
