package test

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var myKey = []byte("go_oj_key")

func TestGenerateToken(t *testing.T) {
	userClaims := &UserClaims{
		Identity:       "u1",
		Username:       "zy",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Error(err)
		return
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InUxIiwidXNlcm5hbWUiOiJ6eSJ9.4F94Wo-MJu33BEU6EpqLrprY2O7CmCyfBzkp8O0WL5U
	t.Log(tokenString)
}

func TestAnalysisToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InUxIiwidXNlcm5hbWUiOiJ6eSJ9.4F94Wo-MJu33BEU6EpqLrprY2O7CmCyfBzkp8O0WL5U"
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Error(err)
		return
	}
	if claims.Valid {
		t.Log(userClaims)
	} else {
		t.Log("token is invalid")
	}
}
