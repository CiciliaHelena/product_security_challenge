package app

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var database *sql.DB

type UserDetails struct {
	username string
	email    string
	password string
}

func NewUserDetails(username, email, password string) *UserDetails {
	return &UserDetails{
		username: username,
		email:    email,
		password: password,
	}
}

func init() {
	var err error
	var statement *sql.Stmt

	database, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS users (username TEXT, email TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

}

func (ud *UserDetails) Store() (err error) {
	statement, err := database.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(ud.username, ud.email, hashPassword(ud.password))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Store to database " + ud.username + ": " + ud.email + " | " + ud.password)

	return
}

func (ud *UserDetails) Authenticate() bool {
	statement, err := database.Prepare("SELECT password FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}

	var hashedPassword string
	err = statement.QueryRow(ud.username).Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	return comparePassword(ud.password, hashedPassword)
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

func comparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true
	}
	return false
}
