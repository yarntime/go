package main

import (
    "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbtemp, err := sql.Open("mysql", "root:password@tcp(192.168.254.44:31786)/test")
	db := dbtemp
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := db.Prepare("create table if not exists dev(id int UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,uid varchar(64),did varchar(64),name varchar(64),qid varchar(64),status char DEFAULT 'u')")
	if stmt != nil {
		stmt.Exec()
		stmt.Close()
	}
	stmt, err = db.Prepare("alter table dev convert to character set utf8 collate utf8_general_ci") 
	if stmt != nil {
		stmt.Exec()
		stmt.Close()
	}

	_ = err
}
