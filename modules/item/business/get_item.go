package business

import (
	"context"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error)
}

type getItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store}
}

func (business *getItemBusiness) GetItem(ctx context.Context, id int) (*model.ToDoItem, error) {
	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}
