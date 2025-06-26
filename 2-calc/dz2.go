package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	display()

	operation := inputOperation()
	returnable := validationOperation(operation)
	if returnable != nil {
		fmt.Printf("Ошибка, %v", returnable)
		return
	}

	stringNumbers := inputNumbers()
	arrayNumbers, err := parseNumbers(stringNumbers)
	if err != nil {
		fmt.Printf("Ошибка, %v", err)
		return
	}
	answer := choosingAnOperation(operation, arrayNumbers)
	output(operation, answer)

}

func display() {
	fmt.Println("--------------------------")
	fmt.Println("    КАЛЬКУЛЯТОР ЧИСЕЛ")
	fmt.Println("--------------------------")
}

func inputOperation() string {
	var operation string
	fmt.Print("Введите операцию(AVG, SUM, MED): ")
	fmt.Scanln(&operation)
	return operation
}

func validationOperation(operation string) error {
	if operation != "AVG" && operation != "SUM" && operation != "MED" {
		return fmt.Errorf("некорректная операция: %s", operation)
	}
	return nil
}

func inputNumbers() string {
	var stringNumbers string
	fmt.Print("Введите числа через запятую(1,2,3,4): ")
	fmt.Scanln(&stringNumbers)
	return stringNumbers
}

func choosingAnOperation(operation string, arrayNumbers []float64) float64 {
	answer := 0.0
	switch operation {
	case "AVG":
		answer = AVGOperation(arrayNumbers)
	case "SUM":
		answer = SUMOperation(arrayNumbers)
	case "MED":
		answer = MEDOperation(arrayNumbers)
	}
	return answer
}

func parseNumbers(stringNumbers string) ([]float64, error) {
	if stringNumbers == "" {
		return nil, fmt.Errorf("пустая строка чисел")
	}

	parts := strings.Split(stringNumbers, ",")
	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		trimmedString := strings.TrimSpace(part)
		if trimmedString == "" {
			continue
		}

		number, err := strconv.ParseFloat(trimmedString, 64)
		if err != nil {
			return nil, fmt.Errorf("некорректное число: %s", trimmedString)
		}
		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("нет чисел для обработки")
	}

	return numbers, nil
}

func AVGOperation(arrayNumbers []float64) float64 {
	sum := 0.0
	for _, num := range arrayNumbers {
		sum = sum + num
	}
	avg := sum / float64(len(arrayNumbers))

	return avg
}

func SUMOperation(arrayNumbers []float64) float64 {
	sum := 0.0
	for _, num := range arrayNumbers {
		sum += num
	}
	return sum
}

func MEDOperation(arrayNumbers []float64) float64 {
	data := make([]float64, len(arrayNumbers))
	copy(data, arrayNumbers)

	sort.Float64s(data)

	n := len(data)
	var median float64

	if n%2 == 1 {
		median = data[n/2]
	} else {
		median = (data[n/2-1] + data[n/2]) / 2
	}

	return median
}

func output(operation string, answer float64) {

	fmt.Printf("%s = %.2f", operation, answer)
}
