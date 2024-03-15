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

	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		// Endpoint получает сотрудников по id компании
		r.Get("/staff/{companyId}", staff.Get)

		// Endpoint добавляет сотрудника
		r.Post("/staff", staff.Add)

		// Endpoint удаляет сотрудника по id
		r.Delete("/staff/{deleteId}", staff.Delete)

		// Endpoint обновляет поля сотрудника по id
		r.Put("/staff", staff.Update)
	})

	return r, nil
}
