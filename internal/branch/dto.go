package branch

type Request struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type DTO struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}
