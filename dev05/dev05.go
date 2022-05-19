/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки

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

type myFlags struct {
	A, B, C, c int
	i, v, F, n bool
}

var MyFlags myFlags

func init() {
	flag.IntVar(&MyFlags.A, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&MyFlags.B, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&MyFlags.C, "C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	flag.IntVar(&MyFlags.c, "c", 0, "'count' (количество строк)")
	flag.BoolVar(&MyFlags.n, "n", false, "line num, напечатать номер строки")
	flag.BoolVar(&MyFlags.i, "i", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&MyFlags.v, "v", false, "'fixed', точное совпадение со строкой, не паттерн")
	flag.BoolVar(&MyFlags.F, "F", false, "'line num', напечатать номер строки")
}

func main() {
	flag.Parse()
	pattern := flag.Arg(0)
	dir, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/dev05/tests/test.txt", dir))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	strs, matchesIdx, out := make([][]string, 0), make([]int, 0), ""
	sc := bufio.NewScanner(file)
	var counter int
	pattern = "ege"
	for sc.Scan() {
		strs = append(strs, strings.Split(sc.Text(), " "))
		if strings.Contains(sc.Text(), pattern) {
			matchesIdx = append(matchesIdx, counter)

			if MyFlags.n {
				out = fmt.Sprintf(out + fmt.Sprintf("%d ", counter))
				counter++
				continue
			} else {
				counter++
				continue
			}
		}
		counter++
	}
	if len(matchesIdx) == 0 {
		return
	}
	grep(strs, matchesIdx, out)
}

func grep(strs [][]string, matchesIdx []int, out string) {
	fmt.Println(matchesIdx)
	for _, v := range matchesIdx {
		fmt.Println(strs[v])
	}
	fmt.Println(out)
}
