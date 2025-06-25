package storage

import (
	"context"

	"github.com/TrienThongLu/GoREST/modules/item/model"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ToDoItemUpdate) error {
	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
