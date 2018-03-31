package main

//testovaci web serveru - jen pro root a jeho index.html

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Web server staruje...")

	http.HandleFunc("/termit/", HandleAllData) //práce s lomítkem - není potřeba

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(".")))) //webserver pro localhost

	http.ListenAndServe(":8080", nil)
}

func HandleAllData(w http.ResponseWriter, r *http.Request) { //vrati vsechna data - celou tabulku
	switch r.Method {
	case "GET":
		fmt.Println("GET ")
	case "OPTIONS":
		HandleOptionsCORS(w, r)
	case "POST":
		fmt.Println("POST ")
	case "PUT":
		fmt.Println("PUT ")
	case "DELETE":
		fmt.Println("DELETE ")
	}
}

func HandleOptionsCORS(w http.ResponseWriter, req *http.Request) {
	//odpověď na volání, pokud by se místo POST klient ptal na OPTIONS
	//tato varianta je pro CORS - https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request
	fmt.Println("OPTIONS ")
	w.Header().Set("Content-Length", "0")
	w.Header().Set("Connection", "keep-alive")
	//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/device")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(200)
	return
}

//kompilace pro EXE aplikaci na Windows

// Windows:      $ GOOS=windows GOARCH=386 go build -v main.exe main.go   //??
// Windows10: 	$ GOOS=windows GOARCH=amd64 go build -v main-64.exe main.go
