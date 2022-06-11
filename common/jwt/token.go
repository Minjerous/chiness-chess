package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Uid  int64
	Type string // REFRESH_TOKEN and TOKEN 用于更新token的标识
	Time time.Time
	jwt.StandardClaims
}
type UserClaims struct {
	Id                          int64
	AccessExpireTime            int64
	RefreshExpireTime           int64
	tokenType, issuer           string
	AccessSecret, RefreshSecret string
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

//  token
func GenToken(claims *UserClaims) (Token, error) {

	accessClaim := MyClaims{
		Uid:  claims.Id,
		Type: "ACCESS_TOKEN",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + claims.AccessExpireTime,
			Issuer:    claims.issuer,
		},
	}

	refreshClaim := MyClaims{
		Uid:  claims.Id,
		Type: "REFRESH_TOKEN",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + claims.AccessExpireTime,
			Issuer:    claims.issuer,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim).SignedString([]byte(claims.AccessSecret))

	if err != nil {
		return Token{}, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim).SignedString([]byte(claims.RefreshSecret))

	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}

func ParseToken(claims *UserClaims, token string) (*MyClaims, bool, error) {

	accessToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(claims.AccessSecret), nil
	})
	if err != nil {
		return nil, false, nil
	}

	//access_token 没有过期
	if claims, ok := accessToken.Claims.(*MyClaims); ok && accessToken.Valid {
		return claims, true, nil
	}
	//
	//refreshToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(claims.RefreshSecret), nil
	//})
	//
	////若果refresh_token 也过期了者要重新登录
	//if err != nil {
	//	return nil, false, err
	//}
	//若果access_token过期判断refresh_token
	//if claims, ok := refreshToken.Claims.(*MyClaims); ok && refreshToken.Valid {
	//	return claims, true, nil
	//}

	return nil, false, errors.New("invalid token")
}
