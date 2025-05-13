package advertise

type UserChecker interface {
	IsUserExists(username string) bool
}
