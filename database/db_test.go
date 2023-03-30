package database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnDb(t *testing.T) {
	db := Instance()
	source := db.Source().(*MySQL)
	conn := source.Conn().(*sql.DB)
	fmt.Println(conn)
}
