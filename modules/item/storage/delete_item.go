package storage

import (
	"context"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

func (sql *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := sql.db.Table(model.ToDoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": model.StatusDeleted,
	}).Error; err != nil {
		return err
	}

	return nil
}
