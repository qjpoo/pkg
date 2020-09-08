package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lib/pq"
)

type Error struct {
	ID        uint64         `db:",omitempty" json:"id"`
	Name      sql.NullString `json:"name"`
	Code      sql.NullInt64  `json:"code"`
	IsSevere  bool           `json:"is_severe"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt pq.NullTime    `json:"deleted_at,omitempty"`
}

type Server struct {
	*sql.DB
}

func main() {
	conn, err := sql.Open("postgres", os.Getenv("TEST_DB"))
	if err != nil {
		log.Fatal(err)
	}
	conn.SetMaxOpenConns(250) // Postgres max is usually ~100
	server := Server{DB: conn}
	http.HandleFunc("/", server.handler)
	addr := ":8000"
	log.Printf("starting server on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (server Server) handler(res http.ResponseWriter, req *http.Request) {
	var codes []Error

	result, err := server.Query("SELECT id, name, code, is_severe, created_at, deleted_at FROM errors;")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()

	for result.Next() {
		var code Error
		if err = result.Scan(&code.ID, &code.Name, &code.Code, &code.IsSevere, &code.CreatedAt, &code.DeletedAt); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		codes = append(codes, code)
	}

	if err = result.Err(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(res).Encode(codes); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}