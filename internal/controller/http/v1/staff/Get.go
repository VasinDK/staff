package staff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Get godoc
// @Summary      Получаем сотрудника
// @Description  Получаем сотрудника по id компании и id департамента
// @Tags         staff
// @Accept       json
// @Produce      plain
// @Param		 department_id		query	  string	true	"id департамента" default(1)
// @Param		 company_id			path	  string	true	"id компании"  default(1)
// @Success      200	{array}		entity.Staff
// @Failure      400	{string}  string	"Некорректный запрос"
// @Failure      404	{string}  string	"Не найдено"
// @Failure      500	{string}  string	"Внутренняя ошибка сервера"
// @Router       /v1/staff/{company_id} [get]
func (staff *Staff) Get(w http.ResponseWriter, r *http.Request) {
	res, err := staff.uc.GetStaff(
		staff.uc.Utils.Atoi(chi.URLParam(r, "companyId")),
		staff.uc.Utils.Atoi(r.URL.Query().Get("department_id")),
	)

	if err != nil {
		staff.l.Error("Path: ", r.URL.Path, err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}
