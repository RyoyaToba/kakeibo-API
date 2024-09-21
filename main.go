package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	// URLのパスパラメータを取得
	vars := mux.Vars(r)
	id := vars["id"]

	// MySQLに接続
	db, err := sql.Open("mysql", "root:rootpassword@tcp(mysql:3306)/testdb")
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var message string
	err = db.QueryRow("SELECT message FROM test WHERE id = ?", id).Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Message not found", http.StatusNotFound)
		} else {
			http.Error(w, "Query error", http.StatusInternalServerError)
		}
		return
	}

	// メッセージを返す
	fmt.Fprintf(w, "Message: %s", message)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/getmessage/{id}", getMessage).Methods("GET")

	// サーバーを8080ポートで起動
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
