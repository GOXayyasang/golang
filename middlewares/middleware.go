package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GOXayyasang/golang/models"
	"github.com/GOXayyasang/golang/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

func VerifyToken(c *fiber.Ctx) error {
	var user models.User
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).SendString("Unauthorize")
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETE_KEY")), nil
	})

	if err != nil {
		return c.Status(401).SendString("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(401).SendString("Invalid token claims")
	}
	userClaim := claims["user"]
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     &user,
		DecodeHook: utils.CustomDecoderHook,
	})
	if err != nil {
		fmt.Print("ERROR : %s", err)
		return c.Status(401).SendString(err.Error())
	}
	err = decoder.Decode(userClaim)
	if err != nil {
		fmt.Print("ERROR2 : %s", err)
		return c.Status(401).SendString(err.Error())
	}
	c.Locals("user", user)
	return c.Next()
}

func GenerateToken(userClaim models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": userClaim,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRETE_KEY")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
