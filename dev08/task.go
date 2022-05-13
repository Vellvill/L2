package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%v, %v\n", os.Stderr, err)
		}
		err = execIn(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func execIn(input string) error {
	input = strings.TrimSuffix(input, "\r\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("No path")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("%v, %v\n", os.Stderr, err)
		}
		_, er := fmt.Println(dir)
		return er
	case "echo":
		for i := 1; i < len(args); i++ {
			fmt.Println(args[i])
		}
		return nil
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
