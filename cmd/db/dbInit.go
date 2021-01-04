package db

import (
	"database/sql"
)

func DbInit(Db *sql.DB) (err error){
	DDLs := []string{CreateUsersAccount, CreatNewsTable, CreatVacansyTable}
	for _, ddl := range DDLs{
		_, err := Db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return
}
