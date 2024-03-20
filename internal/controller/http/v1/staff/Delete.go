package staff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Delete godoc
// @Summary      Удаляет сотрудника
// @Description  Удаляет сотрудника по id
// @Tags         staff
// @Accept       json
// @Produce      plain
// @Param		 id		path	  string	true	"id"
// @Success      200	{string}  string	"Запись удалена"
// @Failure      400	{string}  string	"Некорректный запрос"
// @Failure      404	{string}  string	"Не найдено"
// @Failure      500	{string}  string	"Внутренняя ошибка сервера"
// @Router       /v1/staff/{id} [delete]
func (staff *Staff) Delete(w http.ResponseWriter, r *http.Request) {
	res, err := staff.uc.DelStaffById(staff.uc.Utils.Atoi(chi.URLParam(r, "deleteId")))
	if err != nil {
		staff.l.Error("Path: ", r.URL.Path, err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if res == 0 {
		json.NewEncoder(w).Encode("Запись не найдена")
		return
	}

	json.NewEncoder(w).Encode("Запись удалена")
}
