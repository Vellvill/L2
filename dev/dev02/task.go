package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type myError struct {
	custom string
}

func (e myError) Error() string {
	return e.custom
}

func main() {
	readStdin()
}

func readStdin() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		un, err := Unpack(text)
		if err != nil {
			fmt.Println(err)
		}
		log.Println(un)
	}
}

func Unpack(array string) (string, error) {

	if array == "" {
		return array, nil
	} else if unicode.IsDigit(rune(array[0])) {
		return "", myError{custom: "неверная строка на входе"}
	}

	res := make([]rune, 0)
	for i, v := range []rune(array) {
		j, err := strconv.Atoi(string(v))
		if err == nil {
			for k := 1; k < j; k++ {
				res = append(res, rune(array[i-1]))
			}
		} else {
			res = append(res, v)
		}
	}
	return string(res), nil
}
