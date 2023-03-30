package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL() *MySQL {

	strConnect := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database
	db, err := sql.Open("mysql", strConnect)
	if err != nil {
		log.Println("database / mysql / NewMySQL: ", err)
	}

	return &MySQL{
		conn: db,
	}
}

func (m *MySQL) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func (m *MySQL) Conn() interface{} {
	return m.conn
}
