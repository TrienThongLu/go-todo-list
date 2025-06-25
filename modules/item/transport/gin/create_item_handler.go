package ginitem

import (
	"net/http"

	"github.com/TrienThongLu/GoREST/common"
	"github.com/TrienThongLu/GoREST/modules/item/business"
	"github.com/TrienThongLu/GoREST/modules/item/model"
	"github.com/TrienThongLu/GoREST/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateToDoItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data model.ToDoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStorage(db)
		business := business.NewCreateItemBusiness(store)

		if err := business.CreateNewItem(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
