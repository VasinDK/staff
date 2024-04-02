package integration_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"this_module/config"
	v1 "this_module/internal/controller/http/v1"
	"this_module/internal/infrastructure/repository"
	"this_module/internal/pkg/utils"
	"this_module/internal/usecase"
	"this_module/pkg/logger"
	"this_module/pkg/postgres"

	"github.com/go-chi/chi"
)

var stf *chi.Mux

func TestMain(m *testing.M) {
	// меняем текущий каталог на корневой
	currentDir, _ := os.Getwd()
	parentDir := filepath.Dir(currentDir)
	os.Chdir(parentDir)

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Logger.Level)

	db, err := postgres.New(cfg.Storage)
	if err != nil {
		l.Error("postgres.New", "Error: ", err.Error())
		os.Exit(1)
	}

	staffUseCase := usecase.New(
		utils.New(),
		repository.New(db),
	)

	stf, err = v1.NewRouter(l, staffUseCase)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}

	res := m.Run()

	db.Close()

	os.Exit(res)
}

func TestGet(t *testing.T) {
	recorder := httptest.NewRecorder()

	cases := []struct {
		Address      string
		ResultStatus int
		Result       string
	}{
		{
			Address:      "/v1/staff/1",
			ResultStatus: http.StatusOK,
			Result:       `[{"id":1,"`,
		},
		{
			Address:      "/v1/staff/10000",
			ResultStatus: http.StatusOK,
			Result:       "[]",
		},
		{
			Address:      "/v1/staff/0",
			ResultStatus: http.StatusOK,
			Result:       "[]",
		},
		{
			Address:      "/v1/staff/1?department_id=1",
			ResultStatus: http.StatusOK,
			Result:       `[{"id":1,"`,
		},
	}

	for _, cs := range cases {

		req := httptest.NewRequest(http.MethodGet, cs.Address, nil)

		stf.ServeHTTP(recorder, req)

		if recorder.Code != cs.ResultStatus {
			t.Errorf("Address: %v, Code: %v", cs.Address, recorder.Code)
		}

		if cs.Result == "" {
			return
		}

		if contain := strings.Contains(recorder.Body.String(), cs.Result); contain != true {
			t.Errorf("Address: %v, Body: %v", cs.Address, recorder.Body.String())
		}
	}

}
