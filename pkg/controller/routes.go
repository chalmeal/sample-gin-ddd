package controller

import (
	"os"
	"sample-gin-ddd/pkg/infrastracture/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	{
		store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY_PAIR")))
		r.Use(sessions.Sessions(os.Getenv("SESSION_NAME"), store))
		r.Use(cors.New(config.Cors()))
	}

	api := r.Group("/api")
	{
		appControllers(api)
		adminMasterControllers(api)
		masterControllers(api)
		todoControllers(api)
	}

	return r
}
