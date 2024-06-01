package controller

import (
	"net/http"
	"sample-gin-ddd/pkg/controller/authority"
	"sample-gin-ddd/pkg/controller/middleware"
	"sample-gin-ddd/pkg/controller/response"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func adminMasterControllers(r *gin.RouterGroup) {
	api := r.Group("/admin/master")

	service := usecase.NewMasterAdminService()
	mw := middleware.NewMiddleware()

	// アカウントを検索します。
	{
		api.GET("/account", mw.AuthMiddleware(authority.AUTHORIZE_ADMIN),
			func(c *gin.Context) {
				param := new(model.FindAccount)
				if err := c.BindJSON(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": err.Error(),
					})
					return
				}
				result := service.FindAccount(param)

				response.Res(c, result)
			},
		)
	}
}
