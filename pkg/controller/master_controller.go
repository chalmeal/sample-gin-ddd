package controller

import (
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/infrastracture/session"

	"net/http"
	"sample-gin-ddd/pkg/controller/authority"
	"sample-gin-ddd/pkg/controller/middleware"
	"sample-gin-ddd/pkg/controller/response"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func masterControllers(r *gin.RouterGroup) {
	api := r.Group("/master")

	service := usecase.NewMasterService()
	mw := middleware.NewMiddleware()
	sess := session.NewSessionInfo()

	// 自身のアカウントを取得します。
	{
		api.GET("/account", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				id := sess.GetAccountId(c)
				result := service.GetAccount(id)

				response.Res(c, result)
			},
		)
	}

	// 自身のアカウントを編集します。
	{
		api.PUT("/account/edit", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				param := new(model.EdtAccount)
				if err := c.Bind(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": e.EDIT_ACCOUNT_FAILURE.Error(),
					})
					return
				}

				result := service.EditAccount(sess.GetAccountId(c), param)
				if result.Error != nil {
					response.Res(c, result)
					return
				}

				response.Res(c, result)
			},
		)
	}

	// 自身のアカウントを削除します。
	{
		api.DELETE("/account/delete", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				result := service.DeleteAccount(sess.GetAccountId(c))
				if result.Error != nil {
					response.Res(c, result)
					return
				}
				sess.Clear(c)

				response.Res(c, result)
			},
		)
	}

}
