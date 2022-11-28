package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Username string `json:"username"`
	IsAdmin  int    `json:"is_admin"`
	jwt.StandardClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("go_oj_key")

func GenerateToken(identity, username string, isAdmin int) (string, error) {
	userClaims := &UserClaims{
		Identity:       identity,
		Username:       username,
		IsAdmin:        isAdmin,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalysisToken(tokenString string) (*UserClaims, error) {
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims.Valid {
		return userClaims, nil
	} else {
		return nil, fmt.Errorf("token is invalid")
	}
}

// SendCode
// send verification code to the user's mailbox
func SendCode(toUserEmail, code string) error {
	host := "smtp.qq.com"
	port := "25"
	username := "meet.yuzhang@qq.com"
	password := "ahcgdolgclwohica"

	e := &email.Email{
		From:    username,
		To:      []string{toUserEmail},
		Subject: "Verification code has been sent, please check",
		HTML:    []byte("Your verification code is <h1>" + code + "</h1>"),
	}

	return e.Send(host+":"+port, smtp.PlainAuth("", username, password, host))
}

func GetRandomCode() string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func GetUUID() string {
	return uuid.NewV4().String()
}
