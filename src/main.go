package main

import (
	"fmt"
	"html"
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
		password1: html.EscapeString(r.FormValue("password1")),
		password2: html.EscapeString(r.FormValue("password2")),
	}

	var ok bool
	ok = true
	if err := app.ValidateEmail(form.email); err != nil {
		log.Print(err)
		ok = false
	}
	if err := app.ValidateUsername(form.username); err != nil {
		log.Print(err)
		ok = false
	}
	if err := app.ValidatePassword(form.password1, form.password2); err != nil {
		log.Print(err)
		ok = false
	}

	if ok {
		data := app.NewUserDetails(form.username, form.email, form.password1)
		data.Store()
		fmt.Fprintf(w, "Signup success!")
		return
	}
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
		password: html.EscapeString(r.FormValue("password")),
	}

	if err := app.ValidateUsername(form.username); err != nil {
		log.Print(err)
		tmpl.Execute(w, form)
		return
	}

	data := app.NewUserDetails(form.username, "", form.password)
	if data.Authenticate() {
		fmt.Fprintf(w, "Login success!")
		return
	}

	tmpl.Execute(w, form)
	fmt.Fprintf(w, "Invalid username or password.")
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
