package controllers

import (
	"condo-control/models"
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	jwt "github.com/dgrijalva/jwt-go"
)

// AuthController operations for Middleware
type AuthController struct {
	BaseController
}

//JwtToken =
type JwtToken struct {
	Type    string `json:"type"`
	UserID  string `json:"user_id"`
	CondoID string `json:"condo_id"`
	jwt.StandardClaims
}

//JwtTokenRoute ...
type JwtTokenRoute struct {
	UserID     string              `json:"user_id,omitempty"`
	CondoID    string              `json:"condo_id,omitempty"`
	Points     []*models.Points    `json:"points,omitempty"`
	Assistance *models.Assistances `json:"assistances,omitempty"`
	Item       *models.Items       `json:"item,omitempty"`
	jwt.StandardClaims
}

var hmacSecret = []byte("bazam")

//UserTypes array
var UserTypes = []string{"Admin", "Watcher", "Supervisor", "Resident"}

//VerifyToken =
func VerifyToken(tokenString string, userType string) (decodedToken *JwtToken, err error) {

	if tokenString == "" {
		return nil, errors.New("Empty token")
	}

	tokenString = strings.TrimLeft(tokenString, "Bearer")
	tokenString = strings.TrimLeft(tokenString, " ")

	token, err := jwt.ParseWithClaims(tokenString, &JwtToken{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
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

//VerifyTokenByAllUserTypes ...
func VerifyTokenByAllUserTypes(ts string) (decodedToken *JwtToken, userType string, err error) {

	for _, UserType := range UserTypes {
		userToken, errToken := VerifyToken(ts, UserType)
		if errToken != nil {
			continue
		}
		userType = UserType
		decodedToken = userToken
		return

	}
	err = errors.New("Token is invalid")
	return
}

// GenerateToken =
func (c *BaseController) GenerateToken(userType string, userID string, condoID string, timeArgs ...int) (token string, err error) {

	now := time.Now().In(orm.DefaultTimeLoc)

	timeValues := []int{1, 0, 0}

	for key, timeArg := range timeArgs {
		timeValues[key] = timeArg
	}

	// Create the Claims
	claims := JwtToken{
		userType,
		userID,
		condoID,
		jwt.StandardClaims{
			ExpiresAt: now.AddDate(timeValues[2], timeValues[1], timeValues[0]).Unix(),
			Issuer:    "test",
		},
	}

	var newToken *jwt.Token
	newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = newToken.SignedString(hmacSecret)

	return
}

//GenerateGeneralToken ..
func GenerateGeneralToken(userID string, condoID string, points []*models.Points, assistance *models.Assistances, item *models.Items) (token string, err error) {

	now := time.Now()

	// Create the Claims
	claims := JwtTokenRoute{

		UserID:     userID,
		CondoID:    condoID,
		Points:     points,
		Assistance: assistance,
		Item:       item,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Minute * 30).Unix(),
			Issuer:    "test",
		},
	}

	var newToken *jwt.Token
	newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = newToken.SignedString(hmacSecret)

	return
}

//VerifyGeneralToken ...
func VerifyGeneralToken(tokenString string) (decodedToken *JwtTokenRoute, err error) {

	if tokenString == "" {
		return nil, errors.New("Empty token")
	}

	//tokenString = strings.TrimLeft(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &JwtTokenRoute{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtTokenRoute)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
