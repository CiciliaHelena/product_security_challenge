package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"src/app"
)

type SignupForm struct {
	username  string
	email     string
	password1 string
	password2 string
}

type LoginForm struct {
	username   string
	password   string
	rememberMe bool
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../project/signup.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	form := SignupForm{
		username:  r.FormValue("username"),
		email:     r.FormValue("email"),
		password1: r.FormValue("password1"),
		password2: r.FormValue("password2"),
	}

	log.Println(r.URL.Path, form.username, form.email, form.password1, form.password2)

	if form.password1 != form.password2 {
		fmt.Fprintf(w, "Password rejected!")
		return
	}

	data := app.NewUserDetails(form.username, form.email, form.password1)
	data.Store()

	tmpl.Execute(w, form)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../project/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	form := LoginForm{
		username: r.FormValue("username"),
		password: r.FormValue("password"),
	}

	data := app.NewUserDetails(form.username, "", form.password)
	if data.Authenticate() {
		fmt.Fprintf(w, "Login success!")
		return
	}
	tmpl.Execute(w, form)
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
