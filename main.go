package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", checkHeaders)
	http.HandleFunc("/sql", loadMySQL)

	log.Println("Start server...")
	log.Println(http.ListenAndServe(":8080", nil))
}

func checkHeaders(w http.ResponseWriter, r *http.Request) {
	var result string
	result += fmt.Sprintln("RemoveAddr=", r.RemoteAddr)
	for k, v := range r.Header {
		result += fmt.Sprintln(k, "=", strings.Join(v, ","))
	}
	w.Write([]byte(result))
}

func loadMySQL(w http.ResponseWriter, r *http.Request) {
	var (
		mysqlHost = os.Getenv("DB_HOST")
		mysqlUser = os.Getenv("DB_USER")
		mysqlPass = os.Getenv("DB_PASSWORD")
	)
	diarect := fmt.Sprintf("%s:%s@tcp(%s)/tranz_db", mysqlUser, mysqlPass, mysqlHost)
	db, err := sql.Open("mysql", diarect)
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

	var result string
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
			result += fmt.Sprintln(columns[i], ": ", val)
			result += fmt.Sprintln("-------------------------------------")
		}
	}

	w.Write([]byte(result))
}
