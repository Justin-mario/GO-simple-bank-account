package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var balance float64

func main() {
	message()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("pick a command ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	convertToNumber, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	switch convertToNumber {
	case 1:
		fmt.Println("Enter amount to deposit...")
		amount := bufio.NewReader(os.Stdin)
		payIn, _ := amount.ReadString('\n')
		amountToDeposit, _ := strconv.ParseFloat(strings.TrimSpace(payIn), 64)
		deposit(amountToDeposit)
		fmt.Println("You paid in ", amountToDeposit)
		fmt.Println("You balance is  ", balance)
		main()

	case 2:
		fmt.Println("Enter amount to withdraw...")
		amount := bufio.NewReader(os.Stdin)
		payOut, _ := amount.ReadString('\n')
		amountToWithdraw, _ := strconv.ParseFloat(strings.TrimSpace(payOut), 64)
		withdraw(amountToWithdraw)
		fmt.Println("You withdrew ", amountToWithdraw)
		fmt.Println("You balance is  ", balance)
		main()
	case 3:
		os.Exit(1)

	default:
		fmt.Println("Unknown Command")
		main()
	}
}

func deposit(amount float64) float64 {
	balance += amount
	return balance
}

func withdraw(amount float64) float64 {
	if amount >= 0.0 || balance >= amount {
		balance -= amount
	}
	return balance
}

func message() {
	fmt.Println(` Welcome to Our Bank
		choose a number for your transaction
		1. Deposit.
		2. Withdraw.
		3. Exit. `)

}
