package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

var SECRETRSAPUBLICKEY = []byte("-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMXeLdmBc2phf1Xm8GHd\nEMB9+NZXvR2Qb8P4xFdtUUN6WgeHP01lZ1eYujVevMHCQ2VFVHpAB7vQoMFpUZRL\nSXtAngzX9Zqy0NyZrVUw2cJIV+v90Oqt6rnC07Z71zx4L9NzzVoGfG8dAg/ZMmm2\njGTGv24UttLY/zvhpK3VNBqbnXQvtzrkzsKdDh4jMc5PIZ9aF2DtxFneSAfPrLHt\nEQrd8z1wW6WAqeUe2G/aVZN/ONsFXzXVLj0qWsvB43AvEme0BtlFwz8xq4OL72Pd\nZZML/pSrxsKBu5LZ+0VFfz1TSmoQlb/IKFeVS4msmYk7nJlf5GjX9kn9y06omPCh\nrwIDAQAB\n-----END PUBLIC KEY-----\n")

func JwtProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := strings.Split(c.GetReqHeaders()["Authorization"], "Bearer ")
		if len(authHeader) != 2 {
			logrus.Errorf("Malformed token on request: %s", c.Request().URI())
			return c.Status(http.StatusUnauthorized).JSON("malformed token")
		} else {
			tokenString := authHeader[1]
			isOk, token, err := verifyJWT_RSA(tokenString, SECRETRSAPUBLICKEY)
			if err != nil || !isOk {
				logrus.Errorf("error during verify jwt, err: %s", err.Error())
				return c.Status(http.StatusUnauthorized).JSON(fmt.Sprintf("error during verify jwt, err: %s", err.Error()))
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Context().SetUserValue("tokenClaims", claims)
				return c.Next()
			} else {
				logrus.Errorf("authentication error on request %s: %s", c.Request().URI(), err.Error())
				return c.Status(http.StatusUnauthorized).JSON("unhautorized")
			}

		}
	}
}

// Verify a JWT token using an RSA public key
func verifyJWT_RSA(token string, publicKey []byte) (bool, *jwt.Token, error) {

	var parsedToken *jwt.Token

	// parse token
	state, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		// ensure signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unknown signing method")
		}

		parsedToken = token

		// verify
		key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}

		return key, nil
	})

	if err != nil {
		return false, &jwt.Token{}, err
	}

	if !state.Valid {
		return false, &jwt.Token{}, errors.New("invalid jwt token")
	}

	return true, parsedToken, nil
}

//get current customer from jwt
func GetCustomerNameFromJwt(claims jwt.MapClaims) string {
	return claims["customer_name"].(string)
}

//get current logged in user email (username) from jwt
func GetCurrentLoggedUserEmailFromJwt(claims jwt.MapClaims) string {
	return claims["email"].(string)
}
