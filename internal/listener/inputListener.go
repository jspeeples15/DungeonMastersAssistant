package listener

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Listen() (inputWords []string, err error) {
	// Taking input from user
	in := bufio.NewReader(os.Stdin)

	fullInput, err := in.ReadString('\n')
	fmt.Println(fullInput)
	if err != nil {
		fmt.Sprintf("error taking input: %s, \n", err.Error())
		return []string{"undefined"}, err
	}
	fullInput = strings.TrimSpace(fullInput)
	splitInput := strings.Split(fullInput, " ")
	for i, v := range splitInput {
		//clean it up baby
		splitInput[i] = strings.TrimSpace(v)
	}
	return splitInput, err
}

func Prompt(prompt string) (inputWords []string, err error) {
	fmt.Print(prompt)
	return Listen()
}
