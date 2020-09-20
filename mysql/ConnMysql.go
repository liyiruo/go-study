package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	query()
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func query() {
	db, err := sql.Open("mysql", "root:!QAZ2wsx@tcp(127.0.0.1:3306)/jdbc")
	check(err)
	rows, err := db.Query("select * from myuser s")
	check(err)
	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()
}
