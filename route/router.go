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

		router.POST("user/login", v1.LoginHandler)

		router.POST("user/test", v1.GetParticipateActivityListByUserId)

		loginUserRouter := router.Group("/", middleware.AuthSessionMiddle())
		{
			// 登录用户 参加活动操作
			loginUserRouter.POST("activity/:id/participate", v1.ParticipateActivity)
			loginUserRouter.POST("activity/:id/unparticipate", v1.UnParticipateActivity)

			// 登录用户 个人信息操作
			router.GET("user/userProfile", v1.GetUserProfile)
			router.PUT("user/userProfile", v1.UpdateUserProfile)

			// 登录用户 查看参与的活动
			router.POST("user/test", v1.GetParticipateActivityListByUserId)
		}

		// TODO: 登录用户 评论

		// TODO: 用户查询活动列表  进行的活动 结束的活动
		// TODO: 用户搜索活动 开始时间 结束时间 类型
		// TODO: 用户查询活动详情 参与的用户详情

		router.POST("admin", v1.SaveAdministrator)

		router.POST("admin/login", v1.LoginAdminHandler)

		// TODO: 管理员 登录权限操作
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

		// TODO: 管理员 查询所有用户 用户名 邮箱 头像
	}
	return r
}
