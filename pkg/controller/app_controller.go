package controller

import (
	e "sample-gin-ddd/pkg/errors"

	"net/http"
	"sample-gin-ddd/pkg/controller/response"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func appControllers(r *gin.RouterGroup) {

	service := usecase.NewAppService()

	// アカウントにログインします。
	{
		r.POST("/login", func(c *gin.Context) {
			param := new(model.VerLogin)
			if err := c.Bind(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": e.LOGIN_FAILURE.Error(),
				})
				return
			}

			result, sess := service.Login(param)
			if result.Error == nil {
				sess.Set(c, sess)
			}

			response.Res(c, result)
		})
	}

	// アカウントを仮登録します。
	{
		r.POST("/temporary", func(c *gin.Context) {
			param := new(model.TmpAccount)
			if err := c.Bind(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": e.TEMPORARY_REGISTER_ACCOUNT_FAILURE.Error(),
				})
				return
			}

			result := service.RegisterTemporaryAccount(param)
			response.Res(c, result)
		})
	}

	// アカウントを新規登録します。
	{
		r.POST("",
			func(c *gin.Context) {
				param := new(model.RegAccount)
				if err := c.Bind(&param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": e.REGISTER_ACCOUNT_FAILURE.Error(),
					})
					return
				}

				result := service.RegisterAccount(param)
				if result.Error != nil {
					response.Res(c, result)
					return
				}

				response.Res(c, result)
			},
		)
	}

}
