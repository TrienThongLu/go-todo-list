package storage

import (
	"context"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

func (sql *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error) {
	var data model.ToDoItem

	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
