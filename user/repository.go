package user

type Repository interface {
	Exists(username string) bool
	Save(user User)
}

type InMemoryUserRepository struct {
	users map[string]User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]User),
	}
}

func (r *InMemoryUserRepository) Exists(username string) bool {
	_, exists := r.users[username]
	return exists
}

func (r *InMemoryUserRepository) Save(user User) {
	r.users[user.Username] = user
}
