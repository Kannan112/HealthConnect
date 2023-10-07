package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(TokenString string) (int, error) {
	TokenValue, _ := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("strre"), nil
	})
	var paramsId interface{}
	if claims, ok := TokenValue.Claims.(jwt.MapClaims); ok && TokenValue.Valid {
		paramsId = claims["id"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return 0, fmt.Errorf("token expires")
		}
	}

	value, ok := paramsId.(float64)
	if !ok {
		return 0, fmt.Errorf("expected an int value, but got %T", paramsId)
	}
	id := int(value)
	return id, nil //id
}
