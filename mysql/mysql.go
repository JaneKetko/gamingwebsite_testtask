package mysql

import (
	"database/sql"
	"fmt"

	// Mysql
	_ "github.com/go-sql-driver/mysql"
)

// Session is Mysql session
type Session struct {
	*sql.DB
}

// Open opens connection with dbname  with user,password
func Open(user string, password string, dbname string) (*Session, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return nil, err
	}
	//check that connection is open
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot open DB: %v", err)
	}
	return &Session{db}, nil
}
