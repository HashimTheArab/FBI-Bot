package Utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadToken() string {
	token, err := os.ReadFile("token.tok")
	if err == nil {
		return string(token)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Bot setup!\nToken: ")
	scanner.Scan()
	err = os.WriteFile("token.tok", scanner.Bytes(), 0644)
	if err != nil {
		fmt.Println("Warning: Error saving token file\n" + err.Error())
	}
	return scanner.Text()
}
