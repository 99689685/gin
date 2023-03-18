package middleware

import (
	"ginweb/utils"
	"ginweb/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type Acclaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(4 * time.Hour)
	SetClaims := Acclaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hing",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// CheckToken 验证token
func CheckToken(token string) (*Acclaims, int) {
	betoken, _ := jwt.ParseWithClaims(token, &Acclaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := betoken.Claims.(*Acclaims); code && betoken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ErrorTokenExist

		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong
			c.Abort()
		}
		key, encoded := CheckToken(checkToken[1])
		if encoded == errmsg.ERROR {
			code = errmsg.ErrorTokenWrong
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
