package storage

import (
	"context"

	"github.com/TrienThongLu/GoREST/common"
	"github.com/TrienThongLu/GoREST/modules/item/model"
)

func (sql *sqlStore) GetListItems(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.ToDoItem, error) {
	var result []model.ToDoItem

	db := sql.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if status := f.Status; status != "" {
			db = db.Where("status = ?", status)
		}
	}

	if err := db.Table(model.ToDoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).
		Error; err != nil {
		return nil, err
	}

	return result, nil
}
