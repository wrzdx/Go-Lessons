package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task2/task2/account"
)

func main() {
	account := account.CreateAccount()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		fields := strings.Fields(input)
		if len(fields) == 0 {
			continue
		}
		cmd := fields[0]
		switch cmd {
		case "balance":
			balance, ok := account.Balance()
			if ok != nil {
				fmt.Println(ok.Error())
			} else {

				fmt.Println(balance)
			}
		case "withdraw":
			if len(fields) != 2 {
				fmt.Println("Invalid number of parameters")
				continue
			}
			amount, ok := strconv.Atoi(fields[1])
			if ok != nil {
				fmt.Println(ok.Error())
				continue
			}
			ok = account.Withdraw(amount)
			if ok != nil {
				fmt.Println(ok.Error())
			} else {
				fmt.Println("Successful withdrawal:", amount)
			}
		case "pay":
			if len(fields) != 2 {
				fmt.Println("Invalid number of parameters")
				continue
			}
			amount, ok := strconv.Atoi(fields[1])
			if ok != nil {
				fmt.Println(ok.Error())
				continue
			}
			ok = account.Pay(amount)
			if ok != nil {
				fmt.Println(ok.Error())
			} else {
				fmt.Println("Successful payment:", amount)
			}
			case "topup":
			if len(fields) != 2 {
				fmt.Println("Invalid number of parameters")
				continue
			}
			amount, ok := strconv.Atoi(fields[1])
			if ok != nil {
				fmt.Println(ok.Error())
				continue
			}
			ok = account.TopUp(amount)
			if ok != nil {
				fmt.Println(ok.Error())
			} else {
				fmt.Println("Successful payment:", amount)
			}
		default:
			fmt.Println("Unknown command")
		}
		
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
