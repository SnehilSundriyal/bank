package main

import (
	"example.com/bank/fileOps"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank")
	balance, err := fileOps.GetFloatFromFile(accountBalanceFile, 125000)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Print(err)
		fmt.Print("--------")

	}
	for {
		InitialInstructions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is %f\n", balance)
		} else if choice == 2 {
			var amount float64
			fmt.Print("Enter amount you want to deposit: ")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("Invalid Amount")
				continue
			}

			balance += amount
			fmt.Println("Amount successfully deposited")
			fileOps.WriteFloatToFile(accountBalanceFile, balance)
		} else if choice == 3 {
			var amount float64
			fmt.Print("Enter amount you want to withdraw : ")
			fmt.Scan(&amount)
			if amount < 0 || amount > balance {
				fmt.Println("Invalid Amount")
				continue
			}

			balance -= amount
			fmt.Println("Amount successfully withdrawn")
			fileOps.WriteFloatToFile(accountBalanceFile, balance)
		} else if choice == 4 {
			fmt.Println("Thank You for Visiting!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")
}
