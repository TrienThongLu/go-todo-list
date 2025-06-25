package storage

import (
	"context"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

func (sql *sqlStore) CreateItem(ctx context.Context, data *model.ToDoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
