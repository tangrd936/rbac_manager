package jwts

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"rbac_manager/global"
	"time"
)

type ClaimUserInfo struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	RoleList []uint `json:"role_list"`
}

type Claim struct {
	ClaimUserInfo
	jwt.StandardClaims
}

func GetToken(info ClaimUserInfo) (string, error) {
	j := global.Conf.Jwt
	claims := Claim{info, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(j.Expire) * time.Hour).Unix(), // 过期时间
		Issuer:    j.Issuer,                                                   // 签发人
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret)) // 进行签名生成对应的token
}

func ParseToken(token string) (*Claim, error) {
	j := global.Conf.Jwt
	tk, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := tk.Claims.(*Claim)
	if ok && tk.Valid {
		if claim.Issuer != j.Issuer {
			return nil, errors.New("invalid issuer")
		}
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
