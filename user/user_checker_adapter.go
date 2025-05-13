package user

type CheckerAdapter struct {
	repo Repository
}

func NewUserChecker(repo Repository) *CheckerAdapter {
	return &CheckerAdapter{repo: repo}
}

func (a *CheckerAdapter) IsUserExists(username string) bool {
	return a.repo.Exists(username)
}
