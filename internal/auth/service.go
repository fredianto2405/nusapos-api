package auth

import (
	"errors"
	"github.com/fredianto2405/nusapos-api/pkg/password"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(request *LoginRequest) (*UserDTO, error) {
	user, err := s.repo.FindByUsername(request.Username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	isPasswordMatch := password.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordMatch {
		return nil, errors.New("username atau password salah")
	}

	return user, nil
}
