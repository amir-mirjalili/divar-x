package advertise

type Repository interface {
	Exists(username string) bool
	Save(user Advertise)
}

type InMemoryUserRepository struct {
	advertises map[string]Advertise
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		advertises: make(map[string]Advertise),
	}
}

func (r *InMemoryUserRepository) Exists(title string) bool {
	_, exists := r.advertises[title]
	return exists
}

func (r *InMemoryUserRepository) Save(ads Advertise) {
	r.advertises[ads.Title] = ads
}
