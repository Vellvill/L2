/*
-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	k       int
	n, r, u bool
	path    string
)

type str struct {
	index int
	value string
}

func init() {
	path, _ := os.Getwd()
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.IntVar(&k, "k", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.StringVar(&path, "path", "", "путь к файлу")
}

func main() {
	flag.Parse()
	if len(path) == 0 {
		dir, _ := os.Getwd()
		path = fmt.Sprintf("%s/dev03/tests/test.txt", dir)
	}
	var in io.Reader
	if filename := path; filename == "" {
		fmt.Printf("Не указано имя файла.\n")
		os.Exit(1)
	} else {
		f, err := os.Open(filename)

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Printf("Error closing file: %s", err)
			}
		}(f)

		if err != nil {
			fmt.Printf("Error opening file: %s", err)
			os.Exit(1)
		}
		in = f
	}

	buf := bufio.NewScanner(in)

	strs := make([][]string, 0)

	if u {
		m := make(map[string]struct{})
		for buf.Scan() {
			str := buf.Text()
			if _, ok := m[str]; ok {
				continue
			}
			m[str] = struct{}{}
			strs = append(strs, strings.Split(str, " "))
		}
	} else {
		for buf.Scan() {
			str := buf.Text()
			if len(str)-1 < k {
				continue
			}
			strs = append(strs, strings.Split(str, " "))
		}
	}

	if n {
		sortIndx(strs)
	} else {
		sortAlphabet(strs)
	}
}

func sortIndx(strs [][]string) {
	indexes := make([]int, 0)
	m := make(map[int]string)
	for _, v := range strs {
		a, _ := strconv.Atoi(v[k])
		m[a] = strings.Join(v, " ")
	}
	for _, v := range strs {
		indexes = append(indexes, func() int {
			indx, err := strconv.Atoi(v[k])
			if err != nil {
				log.Println(err)
				return 0
			}
			return indx
		}())
	}
	sort.Ints(indexes)
	if r {
		for i := len(indexes) - 1; i >= 0; i-- {
			if str, ok := m[indexes[i]]; ok {
				fmt.Println(str)
			}
		}
	} else {
		for i := 0; i < len(indexes); i++ {
			if str, ok := m[indexes[i]]; ok {
				fmt.Println(str)
			}
		}
	}
}

func sortAlphabet(strs [][]string) {
	indexes := make([]string, 0)
	m := make(map[string]string)
	for _, v := range strs {
		indexes = append(indexes, v[k])
		m[v[k]] = strings.Join(v, " ")
	}
	sort.Strings(indexes)
	if r {
		for i := len(indexes) - 1; i >= 0; i-- {
			if j, ok := m[indexes[i]]; ok {
				fmt.Println(j)
			}
		}
	} else {
		for _, v := range indexes {
			if j, ok := m[v]; ok {
				fmt.Println(j)
			}
		}
	}
}
