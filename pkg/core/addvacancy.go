package core

import (
	"HumoLab/FinalSolid/cmd/db"
	"HumoLab/FinalSolid/models"
	"database/sql"
	"fmt"
	"log"
)

//Добавляет вакансии
func AddVacancy(database *sql.DB, vacancy models.Vacancy) (ok bool, err error) {
	_, err = database.Exec(db.AddVacancy, vacancy.Name, vacancy.Salary, vacancy.Description, vacancy.DataAdd)
	if err != nil {
		log.Println("Can't insert to News", err)
		return false, err
	}
	fmt.Println(err)
	return true, nil
}


type vacancy struct {
	models.Vacancy
}

//оказывает вакансии
func ShowVacancy(Db *sql.DB) (new []vacancy, err error) {
	rows, err := Db.Query(db.SelShowVacancy)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var p vacancy
		err := rows.Scan(&p.Id, &p.Name, &p.Salary, &p.Description, &p.DataAdd)
		if err != nil {
			log.Println(err)
			continue
		}
		new = append(new, p)
	}
	return
}