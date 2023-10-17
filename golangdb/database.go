package golangdb

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"time"
)

func getCon() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_db")

	if err != nil{
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime( 5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}