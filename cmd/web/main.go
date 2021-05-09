
package main

import (
		"database/sql"
    "log"
    "net/http"
		"flag"
		"os"
		"fmt"
		_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

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
		app.errorLog.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	
	app.infoLog.Println("Database successfully connected!")
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