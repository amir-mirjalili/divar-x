package advertise

import "fmt"

type Repository interface {
	Exists(title string) bool
	Save(ad Advertise)
	Delete(username, title string) error
	FindByTitle(title string) (Advertise, bool)
	GetListByUserName(userName string) []Advertise
}

type InMemoryRepository struct {
	data map[string][]Advertise
}

func NewInMemoryAdRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string][]Advertise),
	}
}

func (r *InMemoryRepository) Exists(title string) bool {
	for _, ads := range r.data {
		for _, ad := range ads {
			if ad.Title == title {
				return true
			}
		}
	}
	return false
}

func (r *InMemoryRepository) Save(ad Advertise) {
	r.data[ad.Username] = append(r.data[ad.Username], ad)
}

func (r *InMemoryRepository) FindByTitle(title string) (Advertise, bool) {
	for _, ads := range r.data {
		for _, ad := range ads {
			if ad.Title == title {
				return ad, true
			}
		}
	}
	return Advertise{}, false
}

func (r *InMemoryRepository) Delete(username, title string) error {
	ads, ok := r.data[username]
	if !ok {
		return fmt.Errorf("access denied")
	}

	var updated []Advertise
	found := false
	for _, ad := range ads {
		if ad.Title == title {
			found = true
			continue
		}
		updated = append(updated, ad)
	}

	if !found {
		return fmt.Errorf("access denied")
	}

	r.data[username] = updated
	return nil
}

func (r *InMemoryRepository) GetListByUserName(username string) []Advertise {
	return r.data[username]
}
