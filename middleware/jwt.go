package middleware

import (
	"errors"
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type JWT struct {
	JwtKey []byte
}
type initilizeJWT struct {
	JwtKey          string
	TokenExpireTime int64
}

var JWTinstance = initilizeJWT{
	JwtKey:          "douyin.com",
	TokenExpireTime: 6000,
}

//创建JWT 实例
func NewJWT() *JWT {
	return &JWT{JwtKey: []byte(JWTinstance.JwtKey)}
}

// 自定义Claim
type tokenClaims struct {
	UserID uint64 //用户ID
	jwt.StandardClaims
}

// SetUpToken 设置 claims，为生成 token 制作准备 "claim"
func SetUpToken(UserID uint64) (string, error) {
	j := NewJWT()
	claims := tokenClaims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 240, //aksdjahskiusaghikasdhfdkajsfghsdkhfjdasakurydgbhkyadxguyesadgwaeyiucftewan
			ExpiresAt: time.Now().Unix() + JWTinstance.TokenExpireTime,
		},
	}

	token, err := j.GenerateToken(claims)
	if err != nil {
		return "", errors.New("token generate error.")
	}
	return token, nil
}

// ParserToken 解析token，返回定义的 Claims
// 如果出现错误，则返回对应的错误信息
func (j *JWT) TranslateToken(tokenString string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JwtKey), nil
	})
	log.Println("tokentranslate----------------")

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token error, please try again.")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token already expired, please try again.")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token invalid, please try again.")
			} else {
				return nil, errors.New("This is not a token, please try again.")
			}
		}

	}
	if token != nil {
		if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token already expired, please try again.")
	}

	return nil, errors.New("This is not a token, please try again.")
}

// CreateToken 通过加密和claim创建token
func (j *JWT) GenerateToken(claims tokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// JWTToken 解析、验证token，并把解析出来的user_id 通过ctx.Set() 方法增加到 gin.Context 头部中
func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//先检查token是否有效
		token := c.Query("token")
		j := NewJWT()
		claims, err := j.TranslateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, commonctl.Response{
				Status_code: -1,
				Status_msg:  err.Error(),
			})
			c.Abort()
			return
		}
		//这个位置加sql查询，通过username和密码 where userid =
		c.Set("middleware_geted_user_id", claims.UserID) //把解析出来的userID放进头部  方便后续逻辑处理
		log.Println(claims.UserID)

	}
}
