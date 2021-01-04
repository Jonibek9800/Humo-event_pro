package token

import (
	"HumoLab/FinalSolid/models"
	"github.com/jwt-go"
	"log"
	"time"
)

type MyCustomClaims struct {
	Id      int64
	Name    string
	Surname string
	Age     int64
	Status  string
	Login   string
	jwt.StandardClaims
}

func CreateToken(user models.Users) string {
	mySigningKey := []byte("AllYourBase")
	claims := MyCustomClaims{
		Id:    user.Id,
		Name:  user.Name,
		Surname: user.Surname,
		Status: user.Status,
		Age: user.Age,
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 36000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)

	}
	return ss
}


func ParseToken(tokenString string) *MyCustomClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	claims := token.Claims.(*MyCustomClaims)
	return claims
}