package integration_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"this_module/config"
	"this_module/internal/controller/http/v1/staff"
	"this_module/internal/infrastructure/repository"
	"this_module/internal/pkg/utils"
	"this_module/internal/usecase"
	"this_module/pkg/logger"
	"this_module/pkg/postgres"
)

var stf staff.Staff

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

	uc := usecase.New(
		utils.New(),
		repository.New(db),
	)

	stf = *staff.New(l, uc)

	res := m.Run()

	db.Close()

	os.Exit(res)
}

func TestGet(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/staff/1", nil)

	stf.Get(recorder, req)

	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(`body: `, body)

	if recorder.Code != http.StatusOK {
		t.Errorf("Get. Recorder.Code: %v", recorder.Code)
	}

}
