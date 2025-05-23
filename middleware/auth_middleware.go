package middleware

import (
	"github.com/gin-gonic/gin"
	"rbac_manager/common"
	black_token "rbac_manager/services/redis_service/token"
	"rbac_manager/utils/jwts"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")

	claims, err := jwts.ParseToken(token)
	if err != nil {
		c.Abort()
		common.FailWithMsg(c, "请登录", err)
		return
	}
	// 判断这个token是否在黑名单中
	if black_token.HaveToken(token) {
		c.Abort()
		common.FailWithMsg(c, "该用户已注销", nil)
		return
	}
	c.Set("claims", claims)
	c.Next()
	return
}

func GetAuth(c *gin.Context) *jwts.Claim {
	claims := c.MustGet("claims")
	return claims.(*jwts.Claim)
}
