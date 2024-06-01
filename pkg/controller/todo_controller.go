package controller

import (
	"net/http"
	"sample-gin-ddd/pkg/controller/authority"
	"sample-gin-ddd/pkg/controller/middleware"
	"sample-gin-ddd/pkg/controller/response"
	"sample-gin-ddd/pkg/infrastracture/session"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func todoControllers(r *gin.RouterGroup) {
	api := r.Group("/todo")

	service := usecase.NewTodoService()
	mw := middleware.NewMiddleware()
	sess := session.NewSessionInfo()

	// 自身のTODOを1つ取得します。
	{
		api.GET("/:id", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				id := c.Param("id")
				accountId := sess.GetAccountId(c)
				result := service.GetTodo(accountId, id)

				response.Res(c, result)
			},
		)
	}

	// 自身のTODO一覧を取得します。
	{
		api.GET("", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				param := new(model.FindTodo)
				if err := c.BindJSON(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": err.Error(),
					})
					return
				}

				accountId := sess.GetAccountId(c)
				result := service.FindTodos(accountId, param)

				response.Res(c, result)
			},
		)
	}

	// 自身のTODOを1つ登録します。
	{
		api.POST("", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				param := new(model.RegTodo)
				if err := c.Bind(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": err.Error(),
					})
					return
				}

				accountId := sess.GetAccountId(c)
				result := service.RegisterTodo(accountId, param)

				response.Res(c, result)
			},
		)
	}

	// 自身のTODOを編集します。
	{
		api.PUT("/edit", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				param := new(model.EdtTodo)
				if err := c.Bind(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": err.Error(),
					})
					return
				}

				accountId := sess.GetAccountId(c)
				result := service.EditTodo(accountId, param)

				response.Res(c, result)
			},
		)
	}

	// 自身のTODOを削除します。
	{
		api.DELETE("/delete", mw.AuthMiddleware(authority.AUTHORIZE_NORMAL),
			func(c *gin.Context) {
				param := new(model.SpecifyTodo)
				if err := c.Bind(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": err.Error(),
					})
					return
				}

				accountId := sess.GetAccountId(c)
				result := service.DeleteTodo(accountId, param.TaskID)

				response.Res(c, result)
			},
		)
	}

}
