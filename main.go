package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var history []Data

type Data struct {
	Num1   float64 `json:"num1"`
	Op     string  `json:"op"`
	Num2   float64 `json:"num2"`
	Result float64 `json:"result"`
}

func getFloatInput(prompt string, reader *bufio.Reader) float64 {

	for {
		fmt.Println(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64) // If conversion successful num save it, if not err will hold an error value
		if err == nil {                           // If err has no error value all good to go.
			return num
		} else {
			fmt.Println("This is an invalid input. Please enter a valid number.")
		}
	}
}

func getCharInput(prompt string, reader *bufio.Reader) string {

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
	fmt.Printf("%.2f %s %.2f = %.2f\n", d.Num1, d.Op, d.Num2, d.Result)
}

func calculations(reader *bufio.Reader) {
	var inputData Data

	inputData.Num1 = getFloatInput("Insert the first number: ", reader)
	inputData.Op = getCharInput("What operation you need? (+, -, /, *)", reader)
	inputData.Num2 = getFloatInput("Write the second number: ", reader)

	switch inputData.Op {
	case "+":
		inputData.Result = inputData.Num1 + inputData.Num2

	case "-":
		inputData.Result = inputData.Num1 - inputData.Num2

	case "*", "x", "X":
		inputData.Result = inputData.Num1 * inputData.Num2

	case "/":
		if inputData.Num2 == 0 {
			fmt.Println("Sorry you can't divide by zero. ")
			return

		} else {
			inputData.Result = inputData.Num1 / inputData.Num2

		}
	default:
		fmt.Printf("Error: invalid input '%s'\n", inputData.Op)
		return
	}

	showResult(inputData)
	history = append(history, inputData)
}

func tempConverter(reader *bufio.Reader) {
	tempIn := getFloatInput("Insert the temperature you have: ", reader)
	unitIn := getCharInput("Insert the unit of the temperature (F, C, K)", reader)
	var tempOut float64
	var unitOut string

	switch unitIn {
	case "F", "f":
		tempOut = (tempIn - 32.0) * 5.0 / 9.0
		unitOut = "ºC"

	case "C", "c":
		tempOut = (tempIn * 9.0 / 5.0) + 32.0
		unitOut = "ºF"

	case "K", "k":
		tempOut = tempIn - 273.15
		unitOut = "ºC"

	default:
		fmt.Println("Sorry, wrong input please use (F, C or K)")
		return
	}

	fmt.Printf("Converted temperature: %.2f %s\n", tempOut, unitOut)
}

func saveHistory() {
	var allHistory []Data // Load existing history to preserve past calculations // Load existing history to preserve past calculations

	fileread, err := os.Open("History.json")

	if err == nil {
		decoder := json.NewDecoder(fileread)
		_ = decoder.Decode(&allHistory)
		fileread.Close()
	}
	// Append current session history
	allHistory = append(allHistory, history...)

	// Write combined history back to file
	file, err := os.Create("History.json")

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t") // Adds identation for readability.
	if err := encoder.Encode(history); err != nil {
		fmt.Println("Error saving history: ", err)
	} else {
		fmt.Println("History save into History.json")
	}
}

func showHistory() {
	file, err := os.Open("History.json")
	if err != nil {
		fmt.Println("No history file found!")
		return
	}
	defer file.Close()
	var loadedHistory []Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&loadedHistory)

	if err != nil {
		fmt.Println("Error reading history:", err)
		return
	}
	fmt.Println("--- History ---")
	for _, d := range loadedHistory {
		showResult(d)
	}
	fmt.Println("------------------")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("What you want to do")
		fmt.Println("1- To use calculator")
		fmt.Println("2- To use a temperature converter")
		fmt.Println("3- To show the history of calculations")
		fmt.Println("4- To exit")

		menu := int(getFloatInput("Insert what you want to do: ", reader))

		switch menu {
		case 1:
			calculations(reader)

		case 2:
			tempConverter(reader)

		case 3:
			showHistory()

		case 4:
			saveHistory()
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Wrong input, please try again!")
		}
	}
}
