package entity

type EntityRepository struct {
	ID         int    `json:"id"`
	Repository string `json:"repository" validate:"required,min=2,max=80"`
	Active     bool   `json:"active"`
}

func (e *EntityRepository) Validate() error {
	return validate.Struct(e)
}

func (e *EntityRepository) GetValidated() error {
	return e.Validate()
}
