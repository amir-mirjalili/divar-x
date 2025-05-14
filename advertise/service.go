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

func (s *Service) Delete(username, title string) (string, error) {
	if !s.userChecker.IsUserExists(username) {
		return "", fmt.Errorf("invalid username")
	}

	ad, found := s.repo.FindByTitle(title)
	if !found {
		return "", fmt.Errorf("invalid title")
	}

	if ad.Username != username {
		return "", fmt.Errorf("access denied")
	}

	err := s.repo.Delete(username, title)
	if err != nil {
		return "", err
	}
	return "removed successfully", nil
}

func (s *Service) GetListByUserName(username string) ([]Advertise, error) {
	if !s.userChecker.IsUserExists(username) {
		return nil, fmt.Errorf("invalid username")
	}
	return s.repo.GetListByUserName(username), nil
}
