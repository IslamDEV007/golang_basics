package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getUserInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created a new bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getUserInput("Choose an option - (a - to add an item, s - to save an item, t - to add a tip) : ", reader)

	switch opt {
	case "a":
		name, _ := getUserInput("item name: ", reader)
		price, _ := getUserInput("item price ($): ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("added an item - ", name, p)

		promptOptions(b)

		fmt.Println(name, price)
	case "t":
		tip, _ := getUserInput("enter the tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", t)

		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)

	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()

	promptOptions(myBill)

}
