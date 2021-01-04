package app

import (
	"HumoLab/FinalSolid/middleware"
	"fmt"
	"log"
	"net/http"

)

func (server *MainServer) InitRoutes () {
	fmt.Println("Routes are init in localhost:8888")

	server.router.GET("/api/humo", server.MainHandler)
	server.router.GET("/api/humo/Sign-in", server.SignInHandler)
	server.router.GET("/api/humo/News",  (middlewares.IsAdmin()(middlewares.Authorized()(server.NewsAll))))
	server.router.GET("/api/humo/Vacancy", middlewares.Authorized()(server.VacancyAll))

	server.router.POST("/api/humo/Vacansy/AddVacancy", middlewares.IsAdmin()(server.AddVacancy))
	server.router.POST("/api/humo/AddNews", (middlewares.IsAdmin()(server.AddNews)))
	server.router.POST("/api/humo/Sign-up", server.SignUpHandler)
	server.router.DELETE("/api/humo/vacancy/delete_vacancy", middlewares.IsAdmin()(server.DeleteVacancy))

	err := http.ListenAndServe("localhost:8888", server)
	if err != nil {
		log.Fatal("Can't listen and serve")
	}
}
