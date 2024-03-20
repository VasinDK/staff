package staff

import (
	"encoding/json"
	"net/http"
	"this_module/internal/entity"
)

// Add godoc
// @Summary      Добавляет сотрудника
// @Description  Добавляет сотрудника
// @Tags         staff
// @Accept       json
// @Produce      plain
// @Param		 request	body	string	true "Сотрудник" SchemaExample({"name": "Кирил", "surname": "Ручников", "phone": "877799977", "companyId": 1, "passportNumber": 1122331, "passportTypeId": 2, "departmentId":1})
// @Success      200	{string}  string	"Ок"
// @Failure      400	{string}  string	"Некорректный запрос"
// @Failure      404	{string}  string	"Не найдено"
// @Failure      500	{string}  string	"Внутренняя ошибка сервера"
// @Router       /v1/staff [post]
func (staff *Staff) Add(w http.ResponseWriter, r *http.Request) {
	StaffExtended := entity.StaffExtended{}

	err := json.NewDecoder(r.Body).Decode(&StaffExtended)
	if err != nil {
		staff.l.Error("Path: ", r.URL.Path, err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	id, err := staff.uc.AddStaff(&StaffExtended)
	if err != nil {
		staff.l.Error("Path: ", r.URL.Path, err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(id)
}
