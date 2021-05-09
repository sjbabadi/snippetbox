
package main

import (
		"database/sql"
    "log"
    "net/http"
		"flag"
		"os"
		"fmt"
		"com.sjbabadi/snippetbox/pkg/models/pg"
		_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *postgres.SnippetModel
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)



	const (
		host = "localhost"
		port = 5432
		user = "sjbabadi"
		dbname = "snippetbox_dev"
	)

	connectionStr := fmt.Sprint("host=%s port=%s user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		errorLog.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		errorLog.Println(err.Error())
	}
	
	infoLog.Println("Database successfully connected!")
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &postgres.SnippetModel{DB: db},
	}
	defer db.Close()

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()


	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}