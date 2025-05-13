package user

import (
	"fmt"
)

type Service struct {
	repo Repository
}

func NewUserService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(username string) error {
	if s.repo.Exists(username) {
		return fmt.Errorf("invalid username")
	}
	user := User{Username: username}
	s.repo.Save(user)
	return nil
}
