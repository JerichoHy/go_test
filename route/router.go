package route

import (
	v1 "awesomeProject/api/v1"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//gin.SetMode(utils.C.Mode)
	r := gin.Default()

	r.Use(middleware.EnableCookieSession())

	router := r.Group("api/v1")
	{
		router.POST("user", v1.SaveUser)
		router.GET("user/:id", v1.GetUser)
		router.GET("users", v1.ListUser)
		router.PUT("user/:id", v1.UpdateUser)
		router.DELETE("user/:id", v1.RemoveUser)

		router.POST("admin", v1.SaveAdministrator)

		router.POST("user/login", v1.LoginHandler)

		router.POST("admin/login", v1.LoginAdminHandler)

		router.POST("activity", v1.SaveActivity)
		router.GET("activity/:id", v1.GetActivity)
		router.GET("activities", v1.ListActivity)
		router.PUT("activity/:id", v1.UpdateActivity)
		router.DELETE("activity/:id", v1.RemoveActivity)

		router.POST("activityCate", v1.SaveActivityCate)
		router.GET("activityCate/:id", v1.GetActivityCate)
		router.GET("activityCates", v1.ListActivityCate)
		router.PUT("activityCate/:id", v1.UpdateActivityCate)
		router.DELETE("activityCate/:id", v1.RemoveActivityCate)

		router.GET("user/userProfile", v1.GetUserProfile)
		router.PUT("user/userProfile", v1.UpdateUserProfile)
	}
	return r
}
