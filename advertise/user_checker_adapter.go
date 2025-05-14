package advertise

type CheckerAdapter struct {
	repo Repository
}

func NewAdsChecker(repo Repository) *CheckerAdapter {
	return &CheckerAdapter{repo: repo}
}

func (a *CheckerAdapter) IsAdsExists(title string) bool {
	return a.repo.Exists(title)
}
