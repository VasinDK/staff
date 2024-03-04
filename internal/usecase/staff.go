// Пакет сценарив использования
package usecase

import (
	"this_module/internal/entity"
)

type Utiler interface {
	Atoi(string) int
}

type Repository interface {
	DelStaffById(int) (int64, error)
	AddStaff(*entity.StaffExtended) (int, error)
	GetStaff(int, int) (*[]entity.Staff, error)
	UpdateStaffById(map[string]any) (bool, error)
}

// Структура с полями репозитория и утилит
type StaffUC struct {
	Utils Utiler
	Repo  Repository
}

func New(Utils Utiler, Repo Repository) *StaffUC {
	return &StaffUC{
		Utils,
		Repo,
	}
}

// Добавляет сотрудника
func (u *StaffUC) AddStaff(staff *entity.StaffExtended) (int, error) {
	id, err := u.Repo.AddStaff(staff)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Получает сотрудника по id компании и id департаментf
func (u *StaffUC) GetStaff(companyId, departmentId int) (*[]entity.Staff, error) {
	staffCompany, err := u.Repo.GetStaff(companyId, departmentId)
	if err != nil {
		return nil, err
	}

	return staffCompany, nil
}

// Удаляет сотрудника по его id
func (u *StaffUC) DelStaffById(id int) (int64, error) {
	res, err := u.Repo.DelStaffById(id)
	if err != nil {
		return 0, err
	}

	return res, nil
}

// Обновляет поля сотрудника по его id. Обновляются только те поля которые пришли в запросе.
// Поле id обязательно, остальные опционально.
func (u *StaffUC) UpdateStaffById(filds map[string]any) (bool, error) {
	res, err := u.Repo.UpdateStaffById(filds)
	if err != nil {
		return false, err
	}

	return res, nil
}
