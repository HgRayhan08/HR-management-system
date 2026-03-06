package domain

import "database/sql"

type EmployeeDomain struct {
	Id           string       `db:"id"`
	UserId       string       `db:"user_id"`
	PositionId   string       `db:"position_id"`
	EmployeeCode string       `db:"employee_code"`
	FullName     string       `db:"last_name"`
	Gender       string       `db:"gender"`
	Address      string       `db:"address"`
	BirthDate    sql.NullTime `db:"birth_date"`
	PhoneNumber  string       `db:"phone_number"`
	IsActive     bool         `db:"is_active"`
	JoinDate     sql.NullTime `db:"join_date"`
	LeaveDate    sql.NullTime `db:"leave_date"`
	Status       string       `db:"status"`
	CreatedAt    sql.NullTime `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}

type EmployeeService interface {
}
