package auth

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByUsername(username string) (*UserDTO, error) {
	sql := `select id, 
		   username, 
		   password_hash as password, 
		   role 
		from nusapos.users 
		where username = $1`

	var user UserDTO
	err := r.db.Get(&user, sql, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
