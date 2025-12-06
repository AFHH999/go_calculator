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

func calculations() {
	var inputData Data

	inputData.num1 = getFloatInput("Insert the first number: ")
	inputData.op = getCharInput(" What operation you need? (+, -, /, *)")
	inputData.num2 = getFloatInput("Write the second number: ")

	switch inputData.op {
	case "+":
		inputData.result = inputData.num1 + inputData.num2

	case "-":
		inputData.result = inputData.num1 - inputData.num2

	case "*":
		inputData.result = inputData.num1 * inputData.num2

	case "/":
		if inputData.num2 == 0 {
			fmt.Println("Sorry you can't divide by zero. ")
			return

		} else {
			inputData.result = inputData.num1 / inputData.num2

		}
	default:
		fmt.Printf("Error: invalid input '%s'\n", inputData.op)
		return
	}

	showResult(inputData)
}

func temp_converter() {
	tempIn := getFloatInput("Insert the temperature you have: ")
	unitIn := getCharInput("Insert the unit of the temperature (F, C, K)")
	var tempOut float64
	var unitOut string

	switch unitIn {
	case "F", "f":
		tempOut = (tempIn - 32.0) * 5.0 / 9.0
		unitOut = "C"

	case "C", "c":
		tempOut = (tempIn * 9.0 / 5.0) + 32.0
		unitOut = "F"

	case "K", "k":
		tempOut = tempIn - 273.15
		unitOut = "F"

	default:
		fmt.Println("Sorry, wrong input please use (F, C or K)")
	}

	fmt.Printf("Converted temperature: %.2f %s\n", tempOut, unitOut)
}

func main() {

	for {
		fmt.Println("What you want to do")
		fmt.Println("1- To use calculator")
		fmt.Println("2- To use a temperature converter")
		fmt.Println("3- To exit")

		menu := int(getFloatInput("Insert what you want to do: "))

		switch menu {
		case 1:
			calculations()

		case 2:
			temp_converter()

		case 3:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Wrong input, please try again!")
		}
	}
}
