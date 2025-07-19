package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var texRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tex Rate: ")
	fmt.Scan(&texRate)

	ebt := revenue - float64(expenses)
	profit := float64(ebt) * (1 - texRate/100)

	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
