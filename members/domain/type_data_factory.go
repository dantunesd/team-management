package domain

import (
	"encoding/json"
	"team-management/members/utils"
)

func TypeDataFactory(typeName Type, data json.RawMessage) (TypeData, error) {
	if typeName == ContractorType {
		return makeContractor(data)
	}

	if typeName == EmployeeType {
		return makeEmployee(data)
	}

	return nil, utils.NewBadRequest("Type must be one of ['employee' 'contractor']")
}

func makeContractor(data json.RawMessage) (*Contractor, error) {
	contractor := &Contractor{}
	return contractor, unmarshal(data, contractor)
}

func makeEmployee(data json.RawMessage) (*Employee, error) {
	employee := &Employee{}
	return employee, unmarshal(data, employee)
}

func unmarshal(data json.RawMessage, output interface{}) error {
	if err := json.Unmarshal(data, output); err != nil {
		return utils.NewBadRequest("The field type_data has invalid fields due to type")
	}
	return nil
}
