package main

import (
	"html/template"
	"log"
	"net/http"
)

type SignupDetails struct {
	username  string
	email     string
	password1 string
	password2 string
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../project/signup.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := SignupDetails{
		username:  r.FormValue("username"),
		email:     r.FormValue("email"),
		password1: r.FormValue("password1"),
		password2: r.FormValue("password2"),
	}

	log.Println(r.URL.Path, details.username, details.email, details.password1, details.password2)

	tmpl.Execute(w, SignupDetails{})
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
