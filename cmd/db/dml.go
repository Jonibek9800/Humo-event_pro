package db

const (
	AddUser = 	`Insert Into Users(name, surname, age, gender, status, login, password) values (($1), ($2), ($3), ($4), ($5), ($6), ($7))`
	AddVacancy = `Insert Into Vacancy(name, salary, description, dataAdd) values (($1), ($2), ($3), ($4))`
	AddNewNews = 	`Insert Into News(name, data, textarea) values (($1), ($2), ($3))`
	SelShowNews  = `select * from News`
	SelShowVacancy  = `select * from Vacancy`
	GetUserByLogIn = `select * from users where login = ($1)`
	DelVacansy = `delete from Vacancy where name = ($1)`
 )