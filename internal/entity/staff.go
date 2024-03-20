// entity пакет с сущностями приложения
package entity

// Сотрудник
type Staff struct {
	Id         int        `json:"id,omitempty" db:"id"`
	Name       string     `json:"name,omitempty" db:"name" example:"Кирил"`
	Surname    string     `json:"surname,omitempty" db:"surname" example:"Ручников"`
	Phone      string     `json:"phone,omitempty" db:"phone" example:"877799977"`
	CompanyId  int        `json:"companyId,omitempty" db:"company_id" example:"2"`
	Passport   Passport   `json:"passport,omitempty" db:"passport"`
	Department Department `json:"department,omitempty" db:"department"`
}

// Расширенная структура - сотрудник
type StaffExtended struct {
	Staff
	PassportNumber int `json:"passportNumber,omitempty" db:"passport_number" example:"111-223"`
	PassportTypeId int `json:"passportTypeId,omitempty" db:"passport_type" example:"2"`
	DepartmentId   int `json:"departmentId,omitempty" db:"department_id" example:"1"`
}

// Паспорт
type Passport struct {
	Type   string `json:"type,omitempty" example:"2"`
	Number string `json:"number,omitempty" example:"112223"`
}

// Департамент
type Department struct {
	Name  string `json:"name,omitempty" example:"бухгалтерия"`
	Phone string `json:"phone,omitempty" example:"22-33-88-888"`
}

// Возвращает map в которой по ключу с названием поля из JSON можно получить название поля в БД
func GetCorrectionMap() map[string]string {
	return map[string]string{
		"companyId":      "company_id",
		"passportNumber": "passport_number",
		"passportTypeId": "passport_type",
		"departmentId":   "department_id",
	}
}
