package repository

import (
	"database/sql"
	"fmt"
	"this_module/internal/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) GetStaff(companyId, departmentId int) (*[]entity.Staff, error) {
	rows := &sql.Rows{}
	err := fmt.Errorf("")

	sql := []byte(`SELECT staff.id, staff.name, surname, staff.phone, company_id, passport_type, 
					passport_number, department.name department_name, department.phone department_phone
				FROM staff 
				LEFT JOIN department 
				ON staff.department_id = department.id`)

	if departmentId > 0 {
		sql = append(sql, '\n')
		sql = append(sql, []byte("WHERE company_id=$1 AND staff.department_id=$2;")...)
		rows, err = r.db.Query(string(sql), companyId, departmentId)
	}

	if departmentId == 0 {
		sql = append(sql, '\n')
		sql = append(sql, []byte("WHERE company_id=$1;")...)
		rows, err = r.db.Query(string(sql), companyId)
	}

	if err != nil {
		return nil, errors.Wrap(err, "r.db.Query")
	}
	defer rows.Close()

	staffCompany := make([]entity.Staff, 0)

	for rows.Next() {
		staff := entity.Staff{}

		err := rows.Scan(&staff.Id, &staff.Name, &staff.Surname, &staff.Phone, &staff.CompanyId,
			&staff.Passport.Type, &staff.Passport.Number, &staff.Department.Name, &staff.Department.Phone)
		if err != nil {
			return nil, errors.Wrap(err, "rows.Next")
		}

		staffCompany = append(staffCompany, staff)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}

	return &staffCompany, nil
}

func (r *Repository) AddStaff(staff *entity.StaffExtended) (int, error) {
	LastInsertId := 0

	sql := `INSERT INTO staff(name, surname, phone, company_id, 
				passport_number, passport_type, department_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	r.db.QueryRow(sql, &staff.Name, &staff.Surname,
		&staff.Phone, &staff.CompanyId, &staff.PassportNumber,
		&staff.PassportTypeId, &staff.DepartmentId).Scan(&LastInsertId)

	return LastInsertId, nil
}

func (r *Repository) DelStaffById(id int) (int64, error) {
	res, err := r.db.Exec(`DELETE FROM staff WHERE id = $1;`, id)
	if err != nil {
		return 0, errors.Wrap(err, "r.db.Exec-DelStaffById")
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "RowsAffected-DelStaffById")
	}

	return RowsAffected, nil
}

func (r *Repository) UpdateStaffById(fields map[string]any) (bool, error) {
	correctionMap := entity.GetCorrectionMap()
	sqlData := make([]any, 0)
	sqlData = append(sqlData, fields["id"])
	sql := []byte("UPDATE staff SET ")
	count := 0

	for i, v := range fields {
		if newName, ok := correctionMap[i]; ok {
			i = newName
		}

		if i != "id" {
			count++
			s := fmt.Sprintf("%v%v", "$", count+1)

			if count > 1 {
				sql = append(sql, ',')
			}

			sql = append(sql, []byte(i)...)
			sql = append(sql, '=')
			sql = append(sql, []byte(s)...)

			sqlData = append(sqlData, v)
		}

	}
	sql = append(sql, []byte(" WHERE id = $1")...)

	res, err := r.db.Exec(string(sql), sqlData...)

	if err != nil {
		return false, errors.Wrap(err, "r.db.Exec-UpdateStaffById")
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "RowsAffected-UpdateStaffById")
	}

	if RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}