package app

import (
	"HumoLab/FinalSolid/cmd/db"
	"HumoLab/FinalSolid/pkg/core/service"
	"net/http"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"log"
)

type MainServer struct {
	Db      *sql.DB
	router  *httprouter.Router
	usersvc *service.UserSvc

}

func NewMainServer(Db *sql.DB, router *httprouter.Router, usersvc *service.UserSvc) *MainServer {
	return &MainServer{Db: Db, router: router, usersvc: usersvc}
}
func (server *MainServer) Start(Db *sql.DB) {
	err := db.DbInit(Db)
	if err != nil {
		log.Fatal("Can't init database err = ", err)
	}
	server.InitRoutes()
}
func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
