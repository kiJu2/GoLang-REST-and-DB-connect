package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	FIRST string `json:"first_name"`
	LAST  string `json:"last_name"`
}

func main() {
	db, err := sql.Open("mysql", "yourId:yourPassword@tcp(yourHost:yourPort(db default is 3306)/schema name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reg")
	// RDS 접근
	defer db.Close()

	fmt.Println("Close")
	// 하나의 Row를 갖는 SQL 쿼리
	start := time.Now()
	results, err := db.Query("SELECT first_name, last_name FROM your_table_name")
	if err != nil {
		log.Fatal(err)
	}
	// 리소스 가져오기
	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.FIRST, &tag.LAST)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	end := time.Now()
	log.Print(start.Nanosecond())
	log.Print(end.Nanosecond())
	log.Print(end.Nanosecond() - start.Nanosecond())
	// 쿼리문을 출력
}
