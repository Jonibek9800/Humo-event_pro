package core

import (
	"database/sql"
	"log"
)

func DeleteVacancy(Db *sql.DB, name string) (err error) {
	_, err = Db.Exec("update Vacancy set Id = ($1) where Name = ($2)", "Removed", name)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
