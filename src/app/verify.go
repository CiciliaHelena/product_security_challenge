package app

import (
	"errors"
	hibp "github.com/mattevans/pwned-passwords"
	"log"
	"regexp"
)

func ValidateEmail(email string) error {
	regex := regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`)
	if !regex.MatchString(email) {
		return errors.New("Email: wrong format")
	}
	if len(email) > 40 {
		return errors.New("Email: exceed 40 characters")
	}
	return nil
}

func ValidateUsername(username string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_.]{8,40}$`)
	if !regex.MatchString(username) {
		return errors.New("Username: wrong format")
	}
	if len(username) < 8 {
		return errors.New("Username: minimum 8 characters")
	}
	if len(username) > 40 {
		return errors.New("Username: exceed 40 characters")
	}
	return nil
}

func ValidatePassword(password1, password2 string) error {
	if password1 != password2 {
		return errors.New("Password mismatch")
	}
	if len(password1) < 8 {
		return errors.New("Password: minimum 10 characters")
	}
	if len(password1) > 40 {
		return errors.New("Password: exceed 100 characters")
	}
	if pwnedPassword(password1) {
		return errors.New("Password: leaked on the internet!")
	}
	return nil
}

func pwnedPassword(password string) bool {
	client := hibp.NewClient()

	pwned, err := client.Pwned.Compromised(password)
	if err != nil {
		log.Print("Pwned check failed")
	}

	return pwned
}
