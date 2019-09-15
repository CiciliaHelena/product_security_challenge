package app

import (
	hibp "github.com/mattevans/pwned-passwords"
	"log"
	"regexp"
)

func ValidateEmail(email string) string {
	regex := regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`)
	if !regex.MatchString(email) {
		return ("Email: wrong format")
	}
	if len(email) > 40 {
		return ("Email: exceed 40 characters")
	}
	return ""
}

func ValidateUsername(username string) string {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_.]{8,40}$`)
	if !regex.MatchString(username) {
		return ("Username: wrong format")
	}
	if len(username) < 8 {
		return ("Username: minimum 8 characters")
	}
	if len(username) > 40 {
		return ("Username: exceed 40 characters")
	}
	return ""
}

func ValidatePassword(password1, password2 string) string {
	if password1 != password2 {
		return ("Password mismatch")
	}
	if len(password1) < 8 {
		return ("Password: minimum 10 characters")
	}
	if len(password1) > 40 {
		return ("Password: exceed 100 characters")
	}
	if pwnedPassword(password1) {
		return ("Password: leaked on the internet!")
	}
	return ""
}

func pwnedPassword(password string) bool {
	client := hibp.NewClient()

	pwned, err := client.Pwned.Compromised(password)
	if err != nil {
		log.Print("Pwned check failed")
	}

	return pwned
}
