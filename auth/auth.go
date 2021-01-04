package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
)

func BasicAuthentication(hand http.HandlerFunc) http.HandlerFunc {
	//GenerateJWT()
	return func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()

		if !ok || username != Username || password != Pass {
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return
		}

		hand.ServeHTTP(w, req)
	}
}

var secretKey = []byte(os.Getenv("SecretKey"))
var Username, Pass string

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	//fmt.Println(token)

	claims := token.Claims.(jwt.MapClaims)
	//claims["authorized"] = true
	//claims["user"] = "admin"
	claims["exp"] = time.Now().Add(300 * time.Second).Unix()

	tokenString, err := token.SignedString(secretKey)
	//fmt.Println("tokenstring: ", tokenString)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

func JwtAuthentication(hand http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {

		tkn := req.Header["Token"][0]

		if len(tkn) == 0 {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Access Denied"))
			return
		}

		token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error Occurs")
			}

			return secretKey, nil
		})

		if err != nil {
			log.Fatal(err)
		}

		if token.Valid {
			hand.ServeHTTP(response, req)
		}
	}
}

func init() {
	err := godotenv.Load("/home/sahadat/go/src/github.com/shahincsejnu/httpAPIserver/.env")
	// err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Username = os.Getenv("UserName")
	Pass = os.Getenv("Password")
}
