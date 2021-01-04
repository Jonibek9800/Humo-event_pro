package service

import (
	"HumoLab/FinalSolid/cmd/db"
	"HumoLab/FinalSolid/cmd/hash"
	"HumoLab/FinalSolid/models"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

type UserSvc struct {
	Db *sql.DB
}
func NewUserSvc(Db *sql.DB) *UserSvc {
	if Db == nil {
		log.Println(errors.New("Db can't be nil"))
	}
	return &UserSvc{Db: Db}
}


func (receiver *UserSvc) RegistUser(Db *sql.DB, user models.SignUp) (err error) {
	Status := "User"
	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		log.Println(http.StatusNetworkAuthenticationRequired, "Can't hash the password")
		return
	}
	_, err = Db.Exec(db.AddUser ,user.Name, user.Surname, user.Age, user.Gender, Status, user.Login, user.Password )
	if err != nil {
		log.Println(err)
		return
	}
	return
}


func (receiver *UserSvc) GetUserByLogin(Db *sql.DB, login string) (user models.Users, err error) {
	row := Db.QueryRow(db.GetUserByLogIn, login)
	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Status,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println(http.StatusAccepted)
		return
	}
	return
}

func (receiver *UserSvc) CheckHasUser(Db *sql.DB, loginBody models.LogIn) (user models.Users, err error) {
	row := Db.QueryRow(db.GetUserByLogIn, loginBody.Login)
	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Status,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println(err)
		return
	}
	if !hash.CheckPasswordHash(loginBody.Password, user.Password) {
		log.Fatal(http.StatusLocked, "invalid password")
		return
	}
	return
}
