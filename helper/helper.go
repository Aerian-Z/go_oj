package helper

import (
	"GOOJ/models"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"os/exec"
	"runtime"
	"sync"
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

func CodeSave(code []byte) (string, error) {
	dirName := "code" + GetUUID()
	path := dirName + "/main.go"
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}
	file, err := os.Create(path)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	if err != nil {
		return "", err
	}

	_, err = file.Write(code)
	if err != nil {
		return "", err
	}
	return path, nil
}

func JudgeCode(pb *models.ProblemBasic, path string) int {
	status := 0
	passCount := 0
	// worry answer
	WA := make(chan int)
	// out of memory
	OOM := make(chan int)
	// compilation error
	CE := make(chan int)
	var lock sync.Mutex

	for _, testCase := range pb.TestCases {
		testCase := testCase
		go func() {
			cmd := exec.Command("go", "run", path)
			var out, stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			stdinPipe, err := cmd.StdinPipe()
			if err != nil {
				log.Fatalln(err)
				return
			}
			_, err = io.WriteString(stdinPipe, testCase.Input)
			if err != nil {
				return
			}
			var bm runtime.MemStats
			runtime.ReadMemStats(&bm)
			if err = cmd.Run(); err != nil {
				log.Println(err, stderr.String())
				if err.Error() == "exit status 2" {
					CE <- 1
					return
				}
			}
			var em runtime.MemStats
			runtime.ReadMemStats(&em)
			if out.String() != testCase.Output {
				WA <- 1
				return
			}
			if (em.Alloc-bm.Alloc)/1024 > uint64(pb.MaxMemory) {
				OOM <- 1
				return
			}

			lock.Lock()
			passCount++
			lock.Unlock()
		}()
	}

	select {
	case <-WA:
		status = 2
	case <-OOM:
		status = 4
	case <-CE:
		status = 5
	case <-time.After(time.Millisecond * time.Duration(pb.MaxRuntime)):
		if passCount == len(pb.TestCases) {
			status = 1
		} else {
			status = 3
		}
	}
	return status
}
