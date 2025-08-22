package branch

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(request *Request) error {
	sql := `insert into nusapos.branches(name, address) 
		values($1, $2)`
	_, err := r.db.Exec(sql, request.Name, request.Address)
	return err
}

func (r *Repository) FindAll(search string) ([]*DTO, error) {
	searchPattern := "%" + search + "%"

	var branches []*DTO
	sql := `select id, name, address 
		from nusapos.branches 
		where name ilike $1
		order by name asc`
	err := r.db.Select(&branches, sql, searchPattern)

	return branches, err
}

func (r *Repository) FindAllPageable(page, size int, search string) ([]*DTO, int, error) {
	searchPattern := "%" + search + "%"
	baseQuery := `where name ilike $1`

	var total int
	countQuery := `select count(0) from nusapos.branches ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	var branches []*DTO
	offset := (page - 1) * size
	dataQuery := `select id, name, address 
		from nusapos.branches 
		` + baseQuery + `
		order by name asc
		limit $2
		offset $3`
	err = r.db.Select(&branches, dataQuery, searchPattern, size, offset)
	if err != nil {
		return nil, 0, err
	}

	return branches, total, nil
}

func (r *Repository) FindByID(id string) (*DTO, error) {
	var branch DTO
	sql := `select id, name, address 
		from nusapos.branches 
		where id = $1
		order by name asc`
	err := r.db.Get(&branch, sql, id)
	return &branch, err
}

func (r *Repository) Update(id string, request *Request) error {
	sql := `update nusapos.branches 
		set name = $1, 
			address = $2 
		where id = $3`
	_, err := r.db.Exec(sql, request.Name, request.Address, id)
	return err
}

func (r *Repository) Delete(id string) error {
	sql := `delete from nusapos.branches where id = $1`
	_, err := r.db.Exec(sql, id)
	return err
}
