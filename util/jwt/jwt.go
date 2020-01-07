package jwt

import (
	"ClockInLite/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(config.ServerSetting.JwtSecret)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

//生成Token
func GenerateToken(Username, Password string, expireDate int64) (string, int64, error) {
	expireTime := time.Now().Unix() + expireDate

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "ClockInLite",
		},
		Username,
		Password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSting, err := token.SignedString(jwtSecret)
	return tokenSting, expireTime, err
}

//解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 刷新token  拿老的token换新的token
//func RefreshToken(token string) (string, time.Time, error) {
//	tokenClaims, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	var expireDate time.Duration = 3600
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			name := claims.Username
//			password := claims.Password
//			return GenerateToken(name, password, expireDate)
//		}
//	}
//
//	return "", time.Now(), nil
//}
