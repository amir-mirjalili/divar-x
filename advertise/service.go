package advertise

import (
	"fmt"
)

type Service struct {
	repo        Repository
	userChecker UserChecker
}

func NewAdsService(repo Repository, userChecker UserChecker) *Service {
	return &Service{repo: repo, userChecker: userChecker}
}

func (s *Service) Insert(username string, title string) error {
	if !s.userChecker.IsUserExists(username) {
		return fmt.Errorf("user is not registered")
	}
	if s.repo.Exists(title) {
		return fmt.Errorf("invalid title")
	}
	advertise := Advertise{Username: username, Title: title}
	s.repo.Save(advertise)
	return nil
}
