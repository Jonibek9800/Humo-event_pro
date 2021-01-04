package main

import (
	"HumoLab/FinalSolid/cmd/app"
	"HumoLab/FinalSolid/pkg/core/service"
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Db, err := sql.Open("sqlite3", "test")
	if err != nil {
		fmt.Println("can not used", err)
	} else {

		router := httprouter.New()
		svc := service.NewUserSvc(Db)
		server := app.NewMainServer(Db, router, svc)
		server.Start(Db)
	}
}