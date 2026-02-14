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

		home()
		fmt.Print("\nChoose a Menu: ")
		var inputMenu int
		fmt.Scan(&inputMenu)

		switch inputMenu {
		case 1:
			// Register
			defer auth.RecoverHandler()
			for {
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
					continue

				}

				result := authService.Register(firstName, lastName, email, pass, confirmPass)
				fmt.Println(result)
				fmt.Println()
				fmt.Println("Enter to continue....")
				fmt.Scanln()
				break
			}
		case 2:
			// Login
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
			fmt.Println()

			if result != "Login successful" {
				fmt.Println("Enter to continue...")
				fmt.Scanln()
				break
			}

		loginMenu:
			for {
				clearscreen.ClearScreen()

				fmt.Println("------ Welcome to system ------")
				fmt.Println()
				fmt.Println("Hello, ", email)
				fmt.Println("1. List Users")
				fmt.Println("2. Logout")
				fmt.Println("\n0. Exit")
				fmt.Print("\nChoose a Menu: ")
				var inputMenu int
				fmt.Scan(&inputMenu)

				switch inputMenu {
				case 1:
					clearscreen.ClearScreen()
					fmt.Println("-------- List Users --------")
					for i, user := range authService.Users {
						fmt.Printf("%d. Full Name: %s %s \n", i+1, user.FirstName, user.LastName)
						fmt.Printf("   Email: %s\n", user.Email)
						fmt.Printf("   Password: %s\n", user.Password)
					}
					fmt.Println()
					fmt.Println("Enter to continue....")
					fmt.Scanln()
				case 2:
					fmt.Println("Logging out...")
					break loginMenu
				case 0:
					fmt.Println("Closing system...")
					os.Exit(0)
				default:
					panic("Input Invalid")
				}
			}

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
			fmt.Println()
		case 0:
			fmt.Println("Closing system...")
			os.Exit(0)
		default:
			panic("Input Invalid")
		}
	}
}
