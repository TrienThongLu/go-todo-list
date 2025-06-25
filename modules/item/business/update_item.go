package business

import (
	"context"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ToDoItemUpdate) error
}

type updateItemBusiness struct {
	store UpdateItemStorage
}

func NewUpdateItemBusiness(store UpdateItemStorage) *updateItemBusiness {
	return &updateItemBusiness{store}
}

func (business *updateItemBusiness) UpdateItem(ctx context.Context, dataUpdate *model.ToDoItemUpdate, id int) error {
	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status == model.StatusDeleted {
		return model.ErrItemDeleted
	}

	if err := business.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
