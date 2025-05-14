package favorite

type UserChecker interface {
	IsUserExists(username string) bool
}
