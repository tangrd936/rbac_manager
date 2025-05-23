package black_token

import (
	"context"
	"rbac_manager/global"
)

func HaveToken(token string) bool {
	_, err := global.Redis.Get(context.Background(), "black_"+token).Result()
	if err != nil {
		return false
	}
	return true
}
