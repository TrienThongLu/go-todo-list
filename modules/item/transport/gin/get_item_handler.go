package ginitem

import (
	"net/http"
	"strconv"

	"github.com/TrienThongLu/GoREST/common"
	"github.com/TrienThongLu/GoREST/modules/item/business"
	"github.com/TrienThongLu/GoREST/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStorage(db)
		business := business.NewGetItemBusiness(store)
		data, err := business.GetItem(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
