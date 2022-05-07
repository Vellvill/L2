package main

import (
	"fmt"
	"log"
	"strconv"
)

type myError struct {
	custom string
}

func (e myError) Error() string {
	return e.custom
}

func main() {
	a, err := Unpack("a1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

func Unpack(array string) (string, error) {

	if array == "" {
		return array, nil
	} else if _, err := strconv.Atoi(string(array[0])); err == nil {
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
