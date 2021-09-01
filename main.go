package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// Conforming to this spec - https://html.spec.whatwg.org/#valid-e-mail-address
var regex = regexp.MustCompile(`^(?P<name>[a-zA-Z0-9.!#$%&'*+/=?^_ \x60{|}~-]+)@(?P<domain>[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*)$`)

func main() {

	for {

		fmt.Println()
		fmt.Println("Enter the email to be validated: ")
		fmt.Println("-------------------------------------------")

		var string string
		fmt.Scanln(&string)

		if isValidEmail(string) {
			color.Set(color.FgGreen)
			fmt.Println("Success, the email address you entered is valid")
			color.Unset()
		} else {
			color.Set(color.FgRed)
			fmt.Println("Invalid email address. Please try again")
			color.Unset()
		}

	}
}

func isValidEmail(s string) bool {

	if len(s) < 3 && len(s) > 254 {
		return false
	}

	if !regex.MatchString(s) {
		return false
	}

	parts := strings.Split(s, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}
