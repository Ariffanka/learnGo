	package golangdb

	import(
		"context"
		"fmt"
		"testing"
		// "database/sql"
		// "golang-mysql/golang_databse"
	)
	
	func TestExecSql(t *testing.T){
		db := getCon()
		defer db.Close()
	
		ctx:= context.Background()
	
		script:= "INSERT INTO test_tb(id,name) VALUES('arip','Arip')"
		_, err:= db.ExecContext(ctx, script)
		if err != nil{
			panic(err)
		}
	
		fmt.Println("Succes add Data")
	}

	func TestQuerySql(t *testing.T){
		db := getCon()
		defer db.Close()
	
		ctx:= context.Background()
	
		script:= "SELECT id,name FROM test_tb"
		rows, err:= db.QueryContext(ctx, script)
		if err != nil{
			panic(err)
		}
	
		defer rows.Close()

		for rows.Next(){
			var id,nama string

			err := rows.Scan(&id,&nama)
			if err!=nil{
				panic(err)
			}

			fmt.Println("id: ", id)
			fmt.Println("nama: ", nama)

		}
	}