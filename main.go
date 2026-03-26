package main

import (
	"fmt"
	"bufio"
	"os"
	"database/sql"

	// import "github.com/go-sql-driver/mysql" -- start here on doc.
)

func userInfoInput(credentials string) string {
	fmt.Print(credentials)
	userInfo := bufio.NewScanner(os.Stdin)
	userInfo.Scan()
	return userInfo.Text()
}

func main(){
	name := userInfoInput("Enter your name: ")
	age := userInfoInput("Enter your age: ")

	fmt.Println("Welcome", name + ".", "You are", age, "years old.")

	// Button handler
	fmt.Println("Press Enter to submit...")
	fmt.Scanln()

}

// User input their information; (name) xx
// Click submit xx
// Save into a DB
// Store inside a UI
// DB display that information