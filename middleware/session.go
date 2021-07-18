package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
