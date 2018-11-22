package main

import(
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)

type DbWorker struct{
	//mysql data source name
	Dsn string
}

func main(){
	dbw := DbWorker{
		Dsn: "root:@tcp(127.0.0.1)/medical",
	}

	db, err := sql.Open("mysql",dbw.Dsn)

	if err != nil{
		panic(err)
		return
	}

	defer db.Close()
}
