package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "dsdddddwdwdwdwddd"
)

//TokenData TokenData
type TokenData struct {
	UserID int64
	jwt.StandardClaims
}

//GetNewToken GetNewToken
func GetNewToken(userID int64) string {
	//根据生成时间，使每次生成的token都不一样
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenData{UserID: userID, StandardClaims: jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
	}})

	jwt.New(jwt.SigningMethodHS256)
	result, err := jwtToken.SignedString([]byte(secretKey))

	if err != nil {
		return err.Error()
	}

	result = base64.StdEncoding.EncodeToString([]byte(result))

	return result
}

//CheckToken CheckToken
func CheckToken(token string) (data interface{}, err error) {
	var bytes []byte
	bytes, err = base64.StdEncoding.DecodeString(token)

	parsedToken, parseErr := jwt.ParseWithClaims(string(bytes), &TokenData{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexcepted signing method:%v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if claims, ok := parsedToken.Claims.(*TokenData); ok && parsedToken.Valid {
		data = claims
		//fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err.Error())
		err = parseErr
		return
	}

	return
}
