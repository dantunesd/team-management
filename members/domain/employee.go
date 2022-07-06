package domain

type Employee struct {
	Role string `json:"role"`
}

func (e *Employee) IsValid() bool {
	return e.Role != ""
}

func (e *Employee) GetType() Type {
	return EmployeeType
}
