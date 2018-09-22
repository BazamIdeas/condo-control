package controllers

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthController operations for Middleware
type AuthController struct {
	BaseController
}

//JwtToken =
type JwtToken struct {
	Type string `json:"tipo"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

var hmacSampleSecret = []byte("bazam")

//VerifyToken =
func VerifyToken(tokenString string, userType string) (decodedToken *JwtToken, err error) {

	if tokenString == "" {
		return nil, errors.New("Token Vacio")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtToken{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtToken)

	if !ok || !token.Valid {
		return nil, err
	}

	//fmt.Println(tokenString, userType)

	if claims.Type != userType {
		return nil, errors.New("Invalid User")
	}

	return claims, nil
}

// GenerateToken =
func (c *BaseController) GenerateToken(userType string, id string, timeArgs ...int) (token string, err error) {

	now := time.Now()

	timeValues := []int{7, 0, 0}

	for key, timeArg := range timeArgs {
		timeValues[key] = timeArg
	}

	// Create the Claims
	claims := JwtToken{
		userType,
		id,
		jwt.StandardClaims{
			ExpiresAt: now.AddDate(timeValues[2], timeValues[1], timeValues[0]).Unix(),
			Issuer:    "test",
		},
	}

	var newToken *jwt.Token
	newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = newToken.SignedString(hmacSampleSecret)

	return
}
