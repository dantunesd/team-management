package domain

type Contractor struct {
	ContractDuration int `json:"contract_duration"`
}

func (c *Contractor) IsValid() bool {
	return c.ContractDuration > 0
}

func (c *Contractor) GetType() Type {
	return ContractorType
}
