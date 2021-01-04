package app

import (
	"HumoLab/FinalSolid/cmd/token"
	"HumoLab/FinalSolid/models"
	"HumoLab/FinalSolid/pkg/core"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (server *MainServer) MainHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("Добро пожаловат в Сайт собитий\n1.Новости\n2.Авторизация\n3.Регистрация"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

//Register now
func (server *MainServer) SignUpHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json: charset = utf-8")
	var requestBody models.SignUp
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	err = server.usersvc.RegistUser(server.Db, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("Этот логин уже занят")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	user, err := server.usersvc.GetUserByLogin(server.Db, requestBody.Login)
	if err != nil {
		log.Println("Can't get new user", err)
		return
	}
	Token := token.CreateToken(user)
	responseToken := models.ResponseToken{
		Description: "Description",
		Token:       Token,
	}
	err = json.NewEncoder(w).Encode(responseToken)
	log.Println("Can't find connection")
	return
}


//Фкнкция входа
func (server *MainServer) SignInHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json: charset = utf-8")
	var requestBody models.LogIn
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	user, err := server.usersvc.CheckHasUser(server.Db, requestBody)
	if err != nil {
		_, err = w.Write([]byte("Вы не зарегистрированы"))
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	Token := token.CreateToken(user)
	responseToken := models.ResponseToken{
		Description: "Description",
		Token:       Token,
	}
	err = json.NewEncoder(w).Encode(responseToken)
	if err != nil {
		log.Println("Can't find connection")
		return
	}
}


//Функция выполняет добавление новастей в База Даных
func (server MainServer) AddNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody models.News
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	_ , err = core.AddNewNews(server.Db, requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write([]byte("Событие добавлено"))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Событие добавлено")
}

//Функция показывает новасти взяв их из База Данных
func (server *MainServer) NewsAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json: charset = utf-8")
	 catalog, err := core.ShowNews(server.Db)
	if err != nil {
		log.Println(err)
		return
	}
	var list []string
	for i := 0; i < len(catalog); i++ {
		list = append(list,  catalog[i].Name, catalog[i].Data, catalog[i].Textarea, " ")
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Println("Can't find connection")
		return
	}
	return
}







func (server MainServer) AddVacancy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody models.Vacancy
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	_ , err = core.AddVacancy(server.Db, requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write([]byte("Вакансия добавлено"))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Вакансия добавлено")
}





//Функция показывает Вакансии взяв их из База Данных
func (server *MainServer) VacancyAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json: charset = utf-8")
	catalog, err := core.ShowVacancy(server.Db)
	if err != nil {
		log.Println(err)
		return
	}
	var list []string
	for i := 0; i < len(catalog); i++ {
		list = append(list,  catalog[i].Name, catalog[i].Salary, catalog[i].Description, catalog[i].DataAdd, " ")
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Println("Can't find connection")
		return
	}
	return
}




func (server *MainServer) DeleteVacancy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	err = core.DeleteVacancy(server.Db, requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write([]byte("Удалено успешно"))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Вакансия удалена")
}
