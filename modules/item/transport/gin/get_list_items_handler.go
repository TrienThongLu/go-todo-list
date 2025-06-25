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

func GetListItems(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		paging.Process()

		var filter model.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStorage(db)
		business := business.NewGetListItemsBusiness(store)

		result, err := business.GetListItems(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
