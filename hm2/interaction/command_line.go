package command_line

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func ReadCommand(welcome_message, command string, allowed_answers map[string]bool) string {
	red := "\033[31m"
	green := "\033[32m"
    reset := "\033[0m"
	purpure := "\033[35m"
	answ := make([]string, 0)
	for v := range allowed_answers {
		answ = append(answ, v)
	}
    fmt.Println(green, welcome_message, reset)
	fmt.Println(purpure, "Available commands: ", command, reset)
	fmt.Println(purpure, "Allowed answers: ", answ, reset)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fields := strings.SplitN(input, " ", 2)
		if len(fields) == 2 {
			if (fields[0] != command) {
				fmt.Println(red, fmt.Sprintf("Invalid command `%v`", fields[0]), reset)
				continue
			}
			_, ok := allowed_answers[fields[1]]
			if (ok) {
				return fields[1]
			} else {
				fmt.Println(red, fmt.Sprintf("Argument `%v` not in allowed answers", fields[1]), reset)
			}
		} else {
			fmt.Println(red, "Invalid input", reset)
			continue
		}
	}
}
