package entity

type Staff struct {
	Id         int        `json:"id,omitempty" db:"id"`
	Name       string     `json:"name,omitempty" db:"name"`
	Surname    string     `json:"surname,omitempty" db:"surname"`
	Phone      string     `json:"phone,omitempty" db:"phone"`
	CompanyId  int        `json:"companyId,omitempty" db:"company_id"`
	Passport   Passport   `json:"passport,omitempty" db:"passport"`
	Department Department `json:"department,omitempty" db:"department"`
}

type StaffExtended struct {
	Staff
	PassportNumber int `json:"passportNumber,omitempty" db:"passport_number"`
	PassportTypeId int `json:"passportTypeId,omitempty" db:"passport_type"`
	DepartmentId   int `json:"departmentId,omitempty" db:"department_id"`
}

type Passport struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type Department struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func GetCorrectionMap() map[string]string {
	return map[string]string{
		"companyId":      "company_id",
		"passportNumber": "passport_number",
		"passportTypeId": "passport_type",
		"departmentId":   "department_id",
	}
}
