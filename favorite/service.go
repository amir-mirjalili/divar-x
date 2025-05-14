package favorite

import "fmt"

type Service struct {
	repo        Repository
	userChecker UserChecker
	adsChecker  AdsChecker
}

func NewFavoriteService(repo Repository, userChecker UserChecker, adsChecker AdsChecker) *Service {
	return &Service{repo: repo, userChecker: userChecker, adsChecker: adsChecker}
}

func (s *Service) Insert(username string, title string) error {
	if !s.userChecker.IsUserExists(username) {
		return fmt.Errorf("invalid username")
	}
	if !s.adsChecker.IsAdsExists(title) {
		return fmt.Errorf("invalid title")
	}
	if s.repo.Exists(title) {
		return fmt.Errorf("already favorite")
	}
	advertise := Favorite{Username: username, Title: title}
	s.repo.Save(advertise)
	return nil
}

func (s *Service) Delete(username, title string) (string, error) {
	if !s.userChecker.IsUserExists(username) {
		return "", fmt.Errorf("invalid username")
	}

	favorite, found := s.repo.FindByTitle(title)
	if !found {
		return "", fmt.Errorf("invalid title")
	}

	if favorite.Username != username {
		return "", fmt.Errorf("access denied")
	}

	err := s.repo.Delete(username, title)
	if err != nil {
		return "", err
	}
	return "removed successfully", nil
}

func (s *Service) GetListByUserName(username string) ([]Favorite, error) {
	if !s.userChecker.IsUserExists(username) {
		return nil, fmt.Errorf("invalid username")
	}
	return s.repo.GetListByUserName(username), nil
}
