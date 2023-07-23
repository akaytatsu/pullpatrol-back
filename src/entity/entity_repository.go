package entity

type EntityRepository struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=80"`
	Url  string `json:"url" validate:"required,min=2,max=455"`
}

func (e *EntityRepository) Validate() error {
	return validate.Struct(e)
}

func (e *EntityRepository) GetValidated() error {
	return e.Validate()
}
