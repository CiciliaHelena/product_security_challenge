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
	Username  string
	Email     string
	password1 string
	password2 string
	Errs      []string
}

type LoginForm struct {
	Username   string
	password   string
	rememberMe bool
	Err        string
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../project/signup.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	form := SignupForm{
		Username:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		password1: html.EscapeString(r.FormValue("password1")),
		password2: html.EscapeString(r.FormValue("password2")),
	}

	var errs []string
	if err := app.ValidateEmail(form.Email); len(err) > 0 {
		log.Print(err)
		errs = append(errs, err)
	}
	if err := app.ValidateUsername(form.Username); len(err) > 0 {
		log.Print(err)
		errs = append(errs, err)
	}
	if err := app.ValidatePassword(form.password1, form.password2); len(err) > 0 {
		log.Print(err)
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		data := app.NewUserDetails(form.Username, form.Email, form.password1)
		data.Store()
		fmt.Fprintf(w, "Signup success! username: "+form.Username+" email: "+form.Email)
		return
	}
	form.Errs = errs
	tmpl.Execute(w, form)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../project/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	form := LoginForm{
		Username: r.FormValue("username"),
		password: html.EscapeString(r.FormValue("password")),
	}

	if err := app.ValidateUsername(form.Username); err == "" {
		log.Printf("Login attempt by " + form.Username)
		data := app.NewUserDetails(form.Username, "", form.password)
		if data.Authenticate() {
			fmt.Fprintf(w, "Login success! username: "+form.Username)
			return
		}
	}

	log.Printf("Failed login attempt by " + form.Username)
	err := "Login failed. Invalid username or password"
	form.Err = err
	tmpl.Execute(w, form)
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
