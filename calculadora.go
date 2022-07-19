package calculator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) parseString(operator string) (int, error) {
	result, err := strconv.Atoi(operator)
	return result, err
}

func (c Calc) Operate(input string, operation string) (int, error) {
	cleanInput := strings.Split(input, operation)
	first, err := c.parseString(cleanInput[0])
	if err != nil {
		return 0, err
	}
	second, err := c.parseString(cleanInput[1])
	if err != nil {
		return 0, err
	}

	switch operation {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		return first / second, nil
	default:
		log.Println(operation, "operation is not supported!")
		return 0, nil

	}

}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ProcessResult(input string, operator string) {
	c := Calc{}
	value, err := c.Operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of", input, "equals to", value)
	}
}
