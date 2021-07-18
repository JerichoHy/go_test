package middleware

import (
	"awesomeProject/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	COOKIE_KEY             = "cookieKey"
	SESSION_NAME           = "seession"
	SESSION_USER_ID_KEY    = "userId"
	SESSION_USER_ROLE_KEY  = "userRole"
	USER_ROLE_USER         = "user"
	USER_ROLE_ADMINSTRATOR = "admin"
)

func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(COOKIE_KEY))
	return sessions.Sessions(SESSION_NAME, store)
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get(SESSION_USER_ID_KEY)
		userRole := session.Get(SESSION_USER_ROLE_KEY)
		if userId == nil || userRole == nil {
			code := utils.NOT_LOGIN
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"data":   userId,
				"msg":    utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if userRole.(string) != USER_ROLE_USER {
			code := utils.SESSION_ERROR
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"data":   userId,
				"msg":    utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}
	}
}
