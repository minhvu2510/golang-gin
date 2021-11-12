package util

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// expireTime := nowTime.Add(3 * time.Hour)
	expireTime := nowTime.Add(1000 * time.Minute)

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	// mySigningKey := []byte("AllYourBase") // for hs256
	// -----------------
	data, err := ioutil.ReadFile("/home/vunm/golang/dev/golang-gin/pkg/util/pri.pem")
	if err != nil {
		fmt.Println(err)
	}

	// If read file success, print content of file
	fmt.Print(string(data))
	key, err := jwt.ParseRSAPrivateKeyFromPEM(data)
	if err != nil {
		fmt.Println("error parsing RSA private key: %v\n", err)
		// return "", fmt.Errorf("error parsing RSA private key: %v\n", err)
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := tokenClaims.SignedString(key)
	// -----------------

	// tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token, err := tokenClaims.SignedString(jwtSecret)
	// token, err := tokenClaims.SignedString(mySigningKey)

	return token, err
}

// // ParseToken parsing token
// func ParseToken(token string) (*Claims, error) {
// 	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		// return jwtSecret, nil
// 		return []byte("AllYourBase"), nil
// 	})

// 	if tokenClaims != nil {
// 		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
// 			return claims, nil
// 		}
// 	}

// 	return nil, err
// }

// ParseToken parsing token RSA 256
func ParseToken(tokenString string) (*jwt.Token, error) {
	publicKey, err := ioutil.ReadFile("/home/vunm/golang/dev/golang-gin/pkg/util/publ.pub")
	if err != nil {
		return nil, fmt.Errorf("error reading public key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA public key: %v\n", err)
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			fmt.Println("----into token method----")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("key---------------------------", key)
		return key, nil
	})
	fmt.Println("---parsedToke:", parsedToken)
	fmt.Println("---err:", err)
	if err != nil {
		fmt.Println("error parsing token: %v", err)
		// return nil, fmt.Errorf("error parsing token: %v", err)
		return nil, err
	}

	return parsedToken, nil
}
