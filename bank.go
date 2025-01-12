package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		panic("Can't continue, sorry.")
		// return
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		fmt.Println("What do you wnat to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
			}
			accountBalance += depositAmount
			fmt.Println("Balance Updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goobye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
			}
			accountBalance += depositAmount
			fmt.Println("Balance Updated! New amount: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Goobye!")
			fmt.Println("Thanks for choosing our bank")
			// return
			// break
		}

	}

}
