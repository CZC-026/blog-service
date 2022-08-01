package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
	//jwt.StandardClaims 结构体，它是 jwt-go 库中预定义的，也是 JWT 的规范，其涉及字段如下：
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//生成token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	//根据claims创建token实例 var1是加密算法 var2是预定义的
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

//解析与校验token
func ParseToken(token string) (*Claims, error) {
	//
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//valid函数验证时间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
