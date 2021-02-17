package main

import "fmt"

var romanNumeralValues map[string]int = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func main() {
	fmt.Printf("Total: %v", romanNumeralToInt("MCMXCIV"))
}

func romanNumeralToInt(romanNumeralString string) int {
	total := 0

	for i := 0; i < len(romanNumeralString); i++ {
		romanNum := string(romanNumeralString[i])

		if (romanNum == "I" || romanNum == "X" || romanNum == "C") && i+1 < len(romanNumeralString) {
			value := romanNumeralValues[romanNum+string(romanNumeralString[i+1])]
			if value > 0 {
				total += value
				i++
				continue
			}
		}

		total += romanNumeralValues[romanNum]
	}

	return total
}
