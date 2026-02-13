package main

import (
	"fmt"
	"koda-b6-golang5/auth"
	clearscreen "koda-b6-golang5/clear-screen"
	"os"
)

func home() {
	defer func() {
		fmt.Println("\n0. Exit")
	}()
	fmt.Println("------ Welcome to system ------")
	fmt.Println()
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Forgot Password")

}
func main() {
	authService := auth.NewAuth()
	for {
		defer func() {

		}()
		home()
		fmt.Print("\nChoose a Menu: ")
		var inputMenu int
		fmt.Scan(&inputMenu)

		switch inputMenu {
		case 1:
			clearscreen.ClearScreen()

			var firstName string
			var lastName string
			var email string
			var pass string
			var confirmPass string
			var confirmData string

			fmt.Println("-------- Register --------")
			fmt.Print("\nWhat is your first name: ")
			fmt.Scanln(&firstName)
			fmt.Print("What is your last name: ")
			fmt.Scanln(&lastName)
			fmt.Print("What is your email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter a strong password: ")
			fmt.Scanln(&pass)
			fmt.Print("Confirm your password: ")
			fmt.Scanln(&confirmPass)

			fmt.Println()
			fmt.Println("Is it true?")
			fmt.Printf("First Name: %s\n", firstName)
			fmt.Printf("Last Name: %s\n", lastName)
			fmt.Printf("Email: %s\n", email)
			fmt.Print("Continue (y/n):")
			fmt.Scan(&confirmData)

			if confirmData != "y" && confirmData != "Y" {
				fmt.Println("Registration cancelled")
				return

			}

			result := authService.Register(firstName, lastName, email, pass, confirmPass)
			fmt.Println(result)
			fmt.Println()
		case 2:
			clearscreen.ClearScreen()

			var email string
			var pass string

			fmt.Println("-------- Login --------")
			fmt.Print("What is your Email: ")
			fmt.Scanln(&email)
			fmt.Print("What is your Password: ")
			fmt.Scanln(&pass)

			result := authService.Login(email, pass)
			fmt.Println(result)

		case 3:
			clearscreen.ClearScreen()

			var email string
			var newPassword string

			fmt.Println("-------- Forgot Password --------")
			fmt.Print("What is your Email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter your new password: ")
			fmt.Scanln(&newPassword)

			result := authService.ForgotPass(email, newPassword)
			fmt.Println(result)
		case 0:
			fmt.Println("Closing system...")
			os.Exit(0)
		default:
			panic("Input Invalid")
		}
	}
}
