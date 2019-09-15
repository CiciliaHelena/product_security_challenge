# Zendesk Product Security Challenge

Implement an authentication mechanism with at least two of the features listed on the [primary README.md](../README.md)

### Prerequisites

make sure to have go installed

## Deployment Setup

Fetch and run the project
```
git clone https://github.com/CiciliaHelena/product_security_challenge
cd product_security_challenge/src
./src
```

## Design Consideration
* Input sanitization and validation
	* Go's html/template package will automatically do data escaping before displaying it to the browser.\
	* The email and username input will be validated using regex.\
		for email; the characters permit are alphanumeric, dot, dash, and underscore which in total does not exceed 40 characters
		for username; the characters permit are alphanumeric, dot, and underscore that consist of 8-40 characters
	* Password can contain any charecter the user want, therefore the string is escaped just after extracting the value
* Password hashed
	Password hasing are implemented by using Provos and Mazières's [bcrypt](golang.org/x/crypto/bcrypt) adaptive hashing algorithm that add salt before the encryption
* Known password check
	Before validate the password on sign up process, the password will be checked if it is already leaked on the internet by calling [HIBP API](https://haveibeenpwned.com/) using a go [client wrapper](https://github.com/mattevans/pwned-passwords)
