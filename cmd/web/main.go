package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	postgres "com.sjbabadi/snippetbox/pkg/models/pg"
	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgres.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	const (
		host     = "localhost"
		port     = "5432"
		user     = "sjbabadi"
		password = ""
		dbname   = "snippetbox_dev"
	)

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	//db, err := sql.Open("postgres", "postgres://sjbabadi:@localhost:5432/snippetbox_dev")
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		//errorLog.Println(err.Error())
		errorLog.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Println("Database successfully connected!")
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgres.SnippetModel{DB: db},
	}
	defer db.Close()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
