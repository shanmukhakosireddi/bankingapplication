// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

var users []User
var currentUser *User

func main() {
	loadUsers()

	for {
		fmt.Println("\n--- Welcome to CLI Bank App ---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			register()
		case 2:
			login()
		case 3:
			saveUsers()
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func register() {
	var u User
	fmt.Print("Enter Name: ")
	fmt.Scan(&u.Name)
	fmt.Print("Enter Age: ")
	fmt.Scan(&u.Age)
	if u.Age < 18 {
		fmt.Println("You must be at least 18 to register.")
		return
	}
	fmt.Print("Enter Email: ")
	fmt.Scan(&u.Email)
	fmt.Print("Enter Phone: ")
	fmt.Scan(&u.Phone)
	fmt.Print("Set Password: ")
	fmt.Scan(&u.Password)
	u.Balance = 0.0
	users = append(users, u)
	saveUsers()
	fmt.Println("Registration successful.")
}

func login() {
	var email, password string
	fmt.Print("Enter Email: ")
	fmt.Scan(&email)
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)

	for i := range users {
		if users[i].Email == email && users[i].Password == password {
			currentUser = &users[i]
			dashboard()
			return
		}
	}
	fmt.Println("Invalid credentials.")
}

func dashboard() {
	for {
		fmt.Printf("\n--- Welcome, %s ---\n", currentUser.Name)
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. View Account Details")
		fmt.Println("5. Update Password")
		fmt.Println("6. Logout")
		fmt.Print("Enter choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Current Balance: ₹%.2f\n", currentUser.Balance)
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			viewDetails()
		case 5:
			updatePassword()
		case 6:
			currentUser = nil
			fmt.Println("Logged out.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func deposit() {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)
	fmt.Print("Confirm deposit (yes/no): ")
	var confirm string
	fmt.Scan(&confirm)
	if confirm == "yes" {
		currentUser.Balance += amount
		saveUsers()
		fmt.Printf("₹%.2f deposited successfully.\n", amount)
	} else {
		fmt.Println("Deposit cancelled.")
	}
}

func withdraw() {
	var amount float64
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)
	if amount > currentUser.Balance {
		fmt.Println("Insufficient balance.")
		return
	}
	currentUser.Balance -= amount
	saveUsers()
	fmt.Printf("₹%.2f withdrawn successfully.\n", amount)
}

func viewDetails() {
	fmt.Printf("\nName: %s\nAge: %d\nEmail: %s\nPhone: %s\nBalance: ₹%.2f\n",
		currentUser.Name, currentUser.Age, currentUser.Email, currentUser.Phone, currentUser.Balance)
}

func updatePassword() {

	var password string

	fmt.Print("Enter Old Password: ")
	fmt.Scan(&password)

	for i := range users {
		if users[i].Password == password {
			currentUser = &users[i]
			var newPass string
			fmt.Print("Enter new password: ")
			fmt.Scan(&newPass)
			currentUser.Password = newPass
			saveUsers()
			fmt.Println("Password updated successfully.")
			return
		}
	}
	fmt.Println("Invalid credentials.")
}

func loadUsers() {
	data, err := ioutil.ReadFile("users.json")
	if err == nil {
		json.Unmarshal(data, &users)
	}
}

func saveUsers() {
	data, _ := json.MarshalIndent(users, "", "  ")
	ioutil.WriteFile("users.json", data, 0644)
}
