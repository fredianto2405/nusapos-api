package category

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(request *Request) error {
	sql := `insert into nusapos.categories(name) values($1)`
	_, err := r.db.Exec(sql, request.Name)
	return err
}

func (r *Repository) FindAll(search string) ([]*DTO, error) {
	searchPattern := "%" + search + "%"

	var categories []*DTO
	sql := `select id, name
		from nusapos.categories 
		where name ilike $1
		order by name asc`
	err := r.db.Select(&categories, sql, searchPattern)

	return categories, err
}

func (r *Repository) FindAllPageable(page, size int, search string) ([]*DTO, int, error) {
	searchPattern := "%" + search + "%"
	baseQuery := `where name ilike $1`

	var total int
	countQuery := `select count(0) from nusapos.categories ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	var categories []*DTO
	offset := (page - 1) * size
	dataQuery := `select id, name 
		from nusapos.categories 
		` + baseQuery + `
		order by name asc
		limit $2
		offset $3`
	err = r.db.Select(&categories, dataQuery, searchPattern, size, offset)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *Repository) FindByID(id string) (*DTO, error) {
	var category DTO
	sql := `select id, name 
		from nusapos.categories 
		where id = $1
		order by name asc`
	err := r.db.Get(&category, sql, id)
	return &category, err
}

func (r *Repository) Update(id string, request *Request) error {
	sql := `update nusapos.categories 
		set name = $1
		where id = $2`
	_, err := r.db.Exec(sql, request.Name, id)
	return err
}

func (r *Repository) Delete(id string) error {
	sql := `delete from nusapos.categories where id = $1`
	_, err := r.db.Exec(sql, id)
	return err
}
