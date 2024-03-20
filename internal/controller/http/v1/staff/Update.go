package staff

import (
	"encoding/json"
	"net/http"
)

// Update godoc
// @Summary      Обновляет данные сотрудника
// @Description  Обновляет данные сотрудника по id
// @Tags         staff
// @Accept       json
// @Produce      plain
// @Param		 request	body	string	true "Сотрудник" SchemaExample({"id": 3, "passportNumber": 99999})
// @Success      200	{string}  string	"Запись обновлена"
// @Failure      400	{string}  string	"Некорректный запрос"
// @Failure      404	{string}  string	"Не найдено"
// @Failure      500	{string}  string	"Внутренняя ошибка сервера"
// @Router       /v1/staff [put]
func (staff *Staff) Update(w http.ResponseWriter, r *http.Request) {
	fields := make(map[string]any)
	json.NewDecoder(r.Body).Decode(&fields)

	res, err := staff.uc.UpdateStaffById(fields)
	if err != nil {
		staff.l.Error("Path: ", r.URL.Path, err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if !res {
		json.NewEncoder(w).Encode("Запись не обновлена")
		return
	}

	json.NewEncoder(w).Encode("Запись обновлена")
}
