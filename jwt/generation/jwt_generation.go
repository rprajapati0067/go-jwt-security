package generation

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userId int64) (string, error) {
	os.Setenv("ACCESS_SECRET", "secret")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute + 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}
