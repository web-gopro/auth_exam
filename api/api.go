package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/api/handlers"
	"github.com/web-gopro/auth_exam/api/middlewares"
	"github.com/web-gopro/auth_exam/redis"
	"github.com/web-gopro/auth_exam/storage"
)

type Options struct {
	Storage storage.StorageI
	Cache   redis.RedisRepoI
}

func Api(o Options) *gin.Engine {

	h := handlers.NewHandlers(o.Cache, o.Storage)

	engine := gin.Default()

	api := engine.Group("/api")
	all := api.Group("/all")

	fmt.Println(h)

	// us.POST("/user", h.UserCreate)

	all.GET("/user/:id", h.GetUserById)
	all.POST("/check", h.CheckUser)
	all.POST("/login", h.Login)
	all.POST("/singup", h.SignUp)

	admp := api.Group("/admp")

	admp.POST("/login", h.SysUserLogin)

	super := api.Group("/super")
	super.Use(middlewares.AuthMiddlewareSuperAdmin())
	super.POST("/singup", h.SysUserSinUp)
	super.GET("/sysuser", h.GetSysUser)


	return engine

}
