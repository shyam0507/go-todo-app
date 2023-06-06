package utility

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(id int) (string, error) {
	tokenExp, _ := strconv.Atoi(os.Getenv("TOKEN_EXP"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Second * time.Duration(tokenExp)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateToken(context *gin.Context) error {
	token, err := parseToken(context)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid")
}

func parseToken(c *gin.Context) (*jwt.Token, error) {
	tokenS := strings.Split(c.Request.Header.Get("Authorization"), " ")
	log.Println(tokenS)
	token, err := jwt.Parse(tokenS[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}
