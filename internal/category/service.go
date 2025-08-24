package category

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(search string) ([]*DTO, error) {
	return s.repo.FindAll(search)
}

func (s *Service) GetAllPageable(page, size int, search string) ([]*DTO, int, error) {
	return s.repo.FindAllPageable(page, size, search)
}

func (s *Service) Create(request *Request) error {
	return s.repo.Insert(request)
}

func (s *Service) Update(id string, request *Request) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return err
	}

	return s.repo.Update(id, request)
}

func (s *Service) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return err
	}

	return s.repo.Delete(id)
}
