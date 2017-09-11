package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", foo)

	log.Println("Start server...")
	log.Println(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func loadMySQL(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tryumph")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer db.Close()

	rows, err := db.Query("select * from Tbl_nouki_check_mst")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			http.Error(w, err.Error(), 500)
		}

		var val string
		for i, col := range values {
			if col == nil {
				val = "NULL"
			} else {
				val = string(col)
			}
			fmt.Println(columns[i], ": ", val)
			fmt.Println("-------------------------------------")
		}
	}

	w.Write([]byte("Success!!"))
}
