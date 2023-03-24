package recipe

import "errors"

type repository struct {
	items []*Model
}

func NewRepository() *repository {
	return &repository{
		items: nil,
	}
}

func (r *repository) Create(rec *Model) (*Model, error) {
	r.items = append(r.items, rec)
	return rec, nil
}

func (r *repository) Read(id uint64) (*Model, error) {
	for _, val := range r.items {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *repository) Update(rec *Model) (*Model, error) {
	var item uint64
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
	list.Recipes = append(list.Recipes, r.items...)
	list.Total = len(r.items)
	return list, nil
}
