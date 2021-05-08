package main

import (
  "log"
  "net/http"
)

// Define a home handler func which writes a byte slice as the res body
func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
	http.NotFound(w, r)
	return
  }

  w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
	w.Header().Set("Allow", http.MethodPost)
	http.Error(w, "Method Not Allowed, 405)
	return
  }
  w.Write([]byte("Create a new snippet..."))
}


func main() {
  // use http.NewServeMux() to initialize a new servemux
  mux := http.NewServeMux()
  //register the home function as the handler for the / URL pattern
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  // use http.ListenAndServe() to start a new web server
	// pass in 2 args - TCP network address to listen on (:4000)
		// and the servemux we just created
	// if ListenAndServe returns an error, log it with log.Fatal() and exit
  log.Println("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
