package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/web-gopro/auth_exam/models"
)

type Claim struct {
	UserId   string
	UserRole string
	jwt.StandardClaims
}

var secretJwtKey = []byte("Shit")

func GenerateJWT(claim models.Claims) (string, error) {

	experationTime := time.Now().Add(1 * time.Hour)

	jwtClaim := Claim{
		UserId:         claim.User_id,
		UserRole:       claim.User_role,
		StandardClaims: jwt.StandardClaims{ExpiresAt: experationTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	tokenString, err := token.SignedString(secretJwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseJWT(tokenString string) (*Claim, error) {

	var claim = &Claim{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {

		return secretJwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {

		return nil, fmt.Errorf("Invalid token")
	}

	return claim, nil

}
