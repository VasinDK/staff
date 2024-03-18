package staff

import (
	"encoding/json"
	"net/http"
	"this_module/internal/entity"
)

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
