package core

import (
	"HumoLab/FinalSolid/cmd/db"
	"HumoLab/FinalSolid/models"
	"database/sql"
	"fmt"
	"log"
)

func AddNewNews(database *sql.DB, news models.News) (ok bool, err error) {
	_, err = database.Exec(db.AddNewNews, news.Name, news.Data, news.Textarea)
	if err != nil {
		log.Println("Can't insert to News", err)
		return false, err
	}
	fmt.Println(err)
	return true, nil
}


type New struct {
	models.News
}
func ShowNews(Db *sql.DB) (new []New, err error) {
	rows, err := Db.Query(db.SelShowNews)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var p New
		err := rows.Scan(&p.Id, &p.Name, &p.Data, &p.Textarea)
		if err != nil {
			log.Println(err)
			continue
		}
		new = append(new, p)
	}
	return
}