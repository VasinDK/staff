package staff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

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
