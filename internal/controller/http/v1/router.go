// Пакет настройки маршрутов
package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"this_module/internal/entity"
	"this_module/internal/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(l *slog.Logger, uc *usecase.StaffUC) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		// Получает сотрудников по id компании. Запрос вида:
		r.Get("/staff/{companyId}", func(w http.ResponseWriter, r *http.Request) {
			res, err := uc.GetStaff(
				uc.Utils.Atoi(chi.URLParam(r, "companyId")),
				uc.Utils.Atoi(r.URL.Query().Get("department_id")),
			)
			if err != nil {
				l.Error("Path: ", r.URL.Path, err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}

			json.NewEncoder(w).Encode(res)
		})

		r.Post("/staff", func(w http.ResponseWriter, r *http.Request) {
			staff := entity.StaffExtended{}

			err := json.NewDecoder(r.Body).Decode(&staff)
			if err != nil {
				l.Error("Path: ", r.URL.Path, err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}

			id, err := uc.AddStaff(&staff)
			if err != nil {
				l.Error("Path: ", r.URL.Path, err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}

			json.NewEncoder(w).Encode(id)
		})

		r.Delete("/staff/{deleteId}", func(w http.ResponseWriter, r *http.Request) {
			res, err := uc.DelStaffById(uc.Utils.Atoi(chi.URLParam(r, "deleteId")))
			if err != nil {
				l.Error("Path: ", r.URL.Path, err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}

			if res == 0 {
				json.NewEncoder(w).Encode("Запись не найдена")
				return
			}

			json.NewEncoder(w).Encode("Запись удалена")
		})

		r.Put("/staff", func(w http.ResponseWriter, r *http.Request) {
			fields := make(map[string]any)
			json.NewDecoder(r.Body).Decode(&fields)

			res, err := uc.UpdateStaffById(fields)
			if err != nil {
				l.Error("Path: ", r.URL.Path, err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}

			if !res {
				json.NewEncoder(w).Encode("Запись не обновлена")
				return
			}

			json.NewEncoder(w).Encode("Запись обновлена")
		})
	})

	return r, nil
}
