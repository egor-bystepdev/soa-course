package command_line

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	chat "hm2/queue_chat/chat"
)


func ReadCommand(welcome_message,
	command string,
	allowed_answers map[string]bool,
	other_commands map[string]func(string),
	stdout_writer *chat.StdoutWriter) string {
	red := "\033[31m"
	green := "\033[32m"
    reset := "\033[0m"
	purpure := "\033[35m"
	answ := make([]string, 0)
	for v := range allowed_answers {
		answ = append(answ, v)
	}
	enter_msg := ""
    enter_msg += fmt.Sprintln(green, welcome_message, reset)
	enter_msg += fmt.Sprintln(purpure, "Available commands: ", command, reset)
	enter_msg += fmt.Sprintln(purpure, "Allowed answers: ", answ, reset)
	enter_msg += purpure + " Also you can enter this commands: "
	for key := range other_commands {
		enter_msg += fmt.Sprint("{", key, "} ")
	}
	enter_msg += reset + "\n"
	stdout_writer.Print(enter_msg)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fields := strings.SplitN(input, " ", 2)
		if len(fields) == 2 {
			if (fields[0] != command) {
				fn, ok := other_commands[fields[0]]
				if !ok {
					stdout_writer.Print(fmt.Sprintln(red, fmt.Sprintf("Invalid command `%v`", fields[0]), reset))
					continue
				}
				fn(fields[1])
				continue
			}
			_, ok := allowed_answers[fields[1]]
			if (ok) {
				return fields[1]
			} else {
				stdout_writer.Print(fmt.Sprintln(red, fmt.Sprintf("Argument `%v` not in allowed answers", fields[1]), reset))
			}
		} else {
			stdout_writer.Print(fmt.Sprintln(red, "Invalid input", reset))
			continue
		}
	}
}
