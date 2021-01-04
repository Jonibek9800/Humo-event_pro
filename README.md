##### Hello everyone who reads my md file
***
***API event is created small part from all available event hope***
***will is evaluated well*** 
***
***here is small part of the code accompaniment vacancies***
```func AddVacancy(database *sql.DB, vacancy models.Vacancy) (ok bool, err error) {
   	_, err = database.Exec(db.AddVacancy, vacancy.Name, vacancy.Salary, vacancy.Description, vacancy.DataAdd)
   	if err != nil {
   		log.Println("Can't insert to News", err)
   		return false, err
   	}
   	fmt.Println(err)
   	return true, nil
   }
```
##Thank uoy for attention

