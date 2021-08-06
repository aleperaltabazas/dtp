package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

func Prompt(s string) string {
	fmt.Printf("%s", s)
	response, err := Reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(response)
}

func PromptConfirmation(s string) bool {
	res := Prompt(fmt.Sprintf("%s (y/n) ", s))
	return Confirm(res)
}

func Confirm(s string) bool {
	if s == "y" {
		return true
	} else if s == "n" {
		return false
	} else {
		return Confirm(Prompt("Please, type 'y' or 'n': "))
	}
}

func GetLine(prefix string) string {
	fmt.Printf(prefix)
	response, err := Reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(response)
}
