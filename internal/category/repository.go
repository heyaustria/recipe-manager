package category

import "errors"

type repository struct {
	items []*Model
}

func NewRepository() *repository {
	return &repository{items: nil}
}

func (r *repository) Create(rec *Model) (*Model, error) {
	r.items = append(r.items, rec)
	return rec, nil
}

func (r *repository) Read(id int64) (*Model, error) {
	for _, val := range r.items {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *repository) Update(rec *Model) (*Model, error) {
	var item int64
	for i, val := range r.items {
		if val.ID == rec.ID {
			r.items[i] = rec
			item = rec.ID
		}
	}
	return r.Read(item)
}

func (r *repository) List() (*List, error) {
	list := &List{}
	list.Items = append(list.Items, r.items...)
	list.Total = len(r.items)
	return list, nil
}

func (r *repository) Exists(cats []*Model) error {
	good := 0
	for _, val := range cats {
		for _, cat := range r.items {
			if val.ID == cat.ID {
				good++
			}
		}
	}
	if good == len(cats) {
		return nil
	} else {
		return errors.New("item not found")
	}
}
