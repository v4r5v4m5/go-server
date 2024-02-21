package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "[-] Error occurred while parsing form - %v", err)
		return
	}
	hackerUsername := r.FormValue("username")
	hackerEmail := r.FormValue("email")
	hackerPassword := r.FormValue("password")
	if hackerUsername == "" || hackerEmail == "" || hackerPassword == "" {
		http.Error(w, "[-] one or more fields are not filled in the form", http.StatusNotAcceptable)
		return
	}
	fmt.Fprintf(w, "[+] Form successfully submitted\n")
	fmt.Fprintf(w, `> h3ll0 fr13nd ! may be i think i should give you a name.
	> h3ll0 "%v"
	> i g0t y0ur addr3ss "%v" 
	> i g0t y0ur passw0rd "%v"`, hackerUsername, hackerEmail, hackerPassword)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { // w - response, r - request
	if r.URL.Path != "/h3ll0" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "> welc0m3 h4ck3r!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // checkout static directory
	http.Handle("/", fileServer)                        // handles / to index.html
	http.HandleFunc("/form", formHandler)               // handles form route with resp func
	http.HandleFunc("/h3ll0", helloHandler)             // handles h3ll0 roue with resp func

	fmt.Printf("[!] Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err) // handles error related to starting web server
	} // creates server and starts listener

}
