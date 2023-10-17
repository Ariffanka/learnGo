package golangdb

import(
	"testing"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)

func TestEmpty(t *testing.T){
	fmt.Println("")
}

func TestCon(t *testing.T){
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_db")

	if err != nil{
		panic(err)
	}
	defer db.Close()

}