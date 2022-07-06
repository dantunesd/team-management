package domain

type Type string

var EmployeeType Type = "employee"
var ContractorType Type = "contractor"

type TypeData interface {
	IsValid() bool
	GetType() Type
}
