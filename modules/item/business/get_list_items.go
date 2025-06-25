package business

import (
	"context"

	"github.com/TrienThongLu/GoREST/common"
	"github.com/TrienThongLu/GoREST/modules/item/model"
)

type GetListItemsStorage interface {
	GetListItems(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.ToDoItem, error)
}

type getListItemsBusiness struct {
	store GetListItemsStorage
}

func NewGetListItemsBusiness(store GetListItemsStorage) *getListItemsBusiness {
	return &getListItemsBusiness{store}
}

func (business *getListItemsBusiness) GetListItems(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.ToDoItem, error) {
	data, err := business.store.GetListItems(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
