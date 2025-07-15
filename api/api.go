package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/api/handlers"
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
	us := api.Group("/us")

	fmt.Println(h)

	// us.POST("/user", h.UserCreate)
	us.GET("/user/:id", h.GetUserById)
	us.POST("/check", h.CheckUser)
	us.POST("/login", h.Login)
	us.POST("/sinup", h.SignUp)


	super:=api.Group("/super")

	super.POST("/create",h.SysUserCreate)

	// us.Use(middlewares.AuthMiddlewareUser())
	// {

	// 	//order
	// 	us.POST("/order", h.CreateOrder)
	// 	us.GET("/order/:id", h.GetOrderById)

	// 	//Author
	// 	us.GET("/auth/:id", h.GetAuthById)

	// 	// book
	// 	us.GET("/book/:id", h.GetBookById)

	// 	// orderItem
	// 	us.POST("/order_item", h.CreateOrderItem)
	// 	us.GET("/order_item/:id", h.GetOrderItemById)
	// 	us.GET("/order_item_id/:id", h.GetOrderItemById)

	// }

	// adm := api.Group("/adm")

	// adm.Use(middlewares.AuthMiddlewareAdmin())
	// {
	// 	// author
	// 	adm.POST("/auth", h.CreateAuth)
	// 	adm.GET("/auth/:id", h.GetAuthById)

	// 	//category
	// 	adm.POST("/category", h.CreateCategory)
	// 	adm.GET("/category/:id", h.GetCategoryById)

	// 	//book
	// 	adm.POST("/book", h.CreateBook)
	// 	adm.GET("/book/:id", h.GetBookById)

	// }

	// all := api.Group("/all")

	// {
	// 	all.GET("/user/:id", h.GetUserById)

	// 	all.POST("/check-user", h.CheckUser) //completed
	// 	all.POST("/sign-up", h.SignUp)       //completed
	// 	all.POST("/sign-in", h.SigIn)        //completed

	// }
	return engine

}
