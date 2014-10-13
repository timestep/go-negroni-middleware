package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	//middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandleFunc) {
	log.Printf("Logging on the wat there...")

	if r.URL.Query().Get("password") == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Printf("Logging on the way back...")
}
