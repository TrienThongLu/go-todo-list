package ginitem

import (
	"net/http"
	"strconv"

	"github.com/TrienThongLu/GoREST/common"
	"github.com/TrienThongLu/GoREST/modules/item/business"
	"github.com/TrienThongLu/GoREST/modules/item/model"
	"github.com/TrienThongLu/GoREST/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data model.ToDoItemUpdate

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStorage(db)
		business := business.NewUpdateItemBusiness(store)

		if err := business.UpdateItem(ctx.Request.Context(), &data, id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
