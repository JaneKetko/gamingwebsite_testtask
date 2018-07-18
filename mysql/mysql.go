package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Session struct {
	*sql.DB
}

func Open(user string, password string, dbname string) (*Session, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return nil, err
	}
	//	defer db.Close()
	//check that connection is open
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Session{db}, nil
}
