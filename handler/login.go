package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const TokenExpireDuration = time.Hour * 4

var MySecret = []byte("ServerSecret")

type MyClaims struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	jwt.StandardClaims
}

func GenToken(username, password string) (string, error) {
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "GoWebTrader",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token.")
}

func Login(c *gin.Context) {

	formData := &MyClaims{}
	msg := ""

	err := c.ShouldBind(formData)
	if err != nil {
		msg = "无效的参数"
		c.JSON(http.StatusOK, gin.H{
			// "status": 2001,
			"msg": msg,
		})
		return
	}

	username, _ := c.Get("username")
	password, _ := c.Get("password")

	if formData.Username == username.(string) && formData.Password == password.(string) {
		tokenString, _ := GenToken(formData.Username, formData.Password)
		msg = "Login success"
		c.JSON(http.StatusOK, gin.H{
			"msg":   msg,
			"token": tokenString,
		})
		log.Println(msg)
		return
	}

	msg = "鉴权失败"
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})

	return
}
