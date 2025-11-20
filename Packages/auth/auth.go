package auth

import "fmt"

// function starting with small letter have package levelscope
// while capital letter have everywhere access

func LoginWithCred(username string, password string) {
	fmt.Println("logging user", username, password)
}
