/*
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	f int
	d string
	s bool
)

type column struct {
	num int
	str string
}

func init() {
	flag.IntVar(&f, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&d, "d", "", "делиметр")
	flag.BoolVar(&s, "s", false, "только строки с разделителем")
}

func main() {
	flag.Parse()
	f = 1
	p, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/dev06/test.txt", p))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		text := sc.Text()
		checker := strings.Split(text, " ")
		if len(checker) < f {
			continue
		}
		fmt.Println(cut(text))
	}
}

func cut(str string) string {
	if s && !strings.Contains(str, d) {
		return ""
	}
	splitted := strings.Split(str, d)
	return splitted[f-1]
}
