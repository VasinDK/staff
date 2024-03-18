package staff

import (
	"encoding/json"
	"net/http"
)

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
