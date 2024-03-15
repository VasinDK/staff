package staff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

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
