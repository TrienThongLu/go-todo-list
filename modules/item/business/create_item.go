package business

import (
	"context"
	"strings"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.ToDoItemCreation) error
}

type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store}
}

func (business *createItemBusiness) CreateNewItem(ctx context.Context, data *model.ToDoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := business.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
