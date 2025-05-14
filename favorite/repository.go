package favorite

import "fmt"

type Repository interface {
	Exists(title string) bool
	Save(favorite Favorite)
	Delete(username, title string) error
	FindByTitle(title string) (Favorite, bool)
	GetListByUserName(userName string) []Favorite
}

type InMemoryRepository struct {
	data map[string][]Favorite
}

func NewInMemoryFavoriteRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string][]Favorite),
	}
}

func (r *InMemoryRepository) Exists(title string) bool {
	for _, favorites := range r.data {
		for _, favorite := range favorites {
			if favorite.Title == title {
				return true
			}
		}
	}
	return false
}

func (r *InMemoryRepository) Save(favorite Favorite) {
	r.data[favorite.Username] = append(r.data[favorite.Username], favorite)
}

func (r *InMemoryRepository) FindByTitle(title string) (Favorite, bool) {
	for _, favorites := range r.data {
		for _, favorite := range favorites {
			if favorite.Title == title {
				return favorite, true
			}
		}
	}
	return Favorite{}, false
}

func (r *InMemoryRepository) Delete(username, title string) error {
	favorites, ok := r.data[username]
	if !ok {
		return fmt.Errorf("access denied")
	}

	var updated []Favorite
	found := false
	for _, favorite := range favorites {
		if favorite.Title == title {
			found = true
			continue
		}
		updated = append(updated, favorite)
	}

	if !found {
		return fmt.Errorf("access denied")
	}

	r.data[username] = updated
	return nil
}

func (r *InMemoryRepository) GetListByUserName(username string) []Favorite {
	return r.data[username]
}
