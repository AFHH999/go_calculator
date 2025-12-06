package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	num1   float64
	op     string
	num2   float64
	result float64
}

func getFloatInput(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return num
		} else {
			fmt.Println("This is an invalid input. Please enter a valid number.")
		}
	}
}

func getCharInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 1 {
			fmt.Println("Invalid input. Please enter exactly one character.")
			continue
		}

		return input
	}
}

func showResult(d Data) {
	fmt.Printf("%.2f %s %.2f = %.2f\n", d.num1, d.op, d.num2, d.result)
}

func main() {

	var input_data Data
	input_data.num1 = getFloatInput("Insert the first number: ")
	input_data.op = getCharInput(" What operation you need? (+, -, /, *)")
	input_data.num2 = getFloatInput("Write the second number: ")

	switch input_data.op {
	case "+":
		input_data.result = input_data.num1 + input_data.num2

	case "-":
		input_data.result = input_data.num1 - input_data.num2

	case "*":
		input_data.result = input_data.num1 * input_data.num2

	case "/":
		if input_data.num2 == 0 {
			fmt.Println("Sorry you can't divide by zero. ")

		} else {
			input_data.result = input_data.num1 / input_data.num2

		}
	}

	showResult(input_data)
}
