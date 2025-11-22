package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	intDivWithAutomaticHandlingAndPrinting(1, 0)
	intDivWithAutomaticHandlingAndPrinting(1, 2)
	intDivWithAutomaticHandlingAndPrinting(4, 2)
}

func intDivWithAutomaticHandlingAndPrinting(numerator, denominator int) {
	if quotient, err := intDiv(numerator, denominator); err != nil {
		fmt.Fprintf(os.Stderr, "integer divide by zero")
	} else {
		fmt.Println(quotient)
	}
}

func intDiv(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("integer divide by zero")
	} else {
		return numerator / denominator, nil
	}
}
