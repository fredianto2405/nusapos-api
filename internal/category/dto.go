package category

type Request struct {
	Name string `json:"name" validate:"required"`
}

type DTO struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
