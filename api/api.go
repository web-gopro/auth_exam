package api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/web-gopro/auth_exam/api/docs"
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
	super.GET("/sysuser", h.GetSysUser)
	super.POST("/sysuser_create", h.SysUserCreate)
	super.POST("/role", h.RoleCreate)
	super.PUT("/role", h.RoleCreate)
	super.GET("/role/:id", h.RoleGetById)


	url := ginSwagger.URL("swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return engine

}
