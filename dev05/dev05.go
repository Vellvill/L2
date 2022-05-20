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
	A, B, C       int
	i, v, F, n, c bool
}

var MyFlags myFlags

func init() {
	flag.IntVar(&MyFlags.A, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&MyFlags.B, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&MyFlags.C, "C", 0, "'context' печатать ±N строк вокруг совпадения")
	flag.BoolVar(&MyFlags.c, "c", false, "'count' (количество строк)")
	flag.BoolVar(&MyFlags.n, "n", false, "line num, напечатать номер строки")
	flag.BoolVar(&MyFlags.i, "i", false, "игнорировать регистр")
	flag.BoolVar(&MyFlags.v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&MyFlags.F, "F", false, "точное совпадение со строкой, не паттерн")
}

func main() {

	flag.Parse()

	pattern := flag.Args()[0]
	filepath := flag.Args()[1]

	dir, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/dev05/%s", dir, filepath))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	strs, matchesIdx := make([][]string, 0), make([]int, 0)
	var out string

	sc := bufio.NewScanner(file)
	var counter int
	for sc.Scan() {
		line := sc.Text()

		if MyFlags.i {
			line = strings.ToLower(line)
		}

		if strings.Contains(line, pattern) {
			matchesIdx = append(matchesIdx, counter)

			if MyFlags.n {
				out = fmt.Sprintf(out + fmt.Sprintf("Matched row number: %d\n", counter+1))
			}

			if MyFlags.v {
				continue
			}

			strs = append(strs, strings.Split(line, " "))
			counter++

		} else {
			strs = append(strs, strings.Split(line, " "))
			counter++
		}
	}
	if len(matchesIdx) == 0 {
		return
	}
	grep(strs, matchesIdx, out)
}

func grep(strs [][]string, matchesIdx []int, out string) {
	var grepABCm map[string][][]string
	if MyFlags.A != 0 || MyFlags.B != 0 || MyFlags.C != 0 {
		grepABCm = grepABC(strs, matchesIdx)
		for k, v := range grepABCm {
			fmt.Println(k)
			for _, raw := range v {
				fmt.Println(raw)
			}
		}
	}
	if MyFlags.c {
		out = fmt.Sprintf(out + fmt.Sprintf("Number of matched rows: %d", len(matchesIdx)))
	}
	if MyFlags.v {
		for _, v := range strs {
			fmt.Println(v)
		}
	} else {
		for _, v := range matchesIdx {
			fmt.Println(strs[v])
		}
	}
	fmt.Println(out)
}

func grepABC(strs [][]string, matchesIdx []int) map[string][][]string {
	res := make(map[string][][]string)
	if MyFlags.A != 0 || MyFlags.B != 0 {
		for _, v := range matchesIdx {
			if MyFlags.B != 0 {
				for i, j := v-1, 0; i >= 0; i, j = i-1, j+1 {
					if j > MyFlags.B-1 {
						break
					}
					res["GREP+BEFORE:"+strings.Join(strs[v], " ")] = append(res["GREP+BEFORE:"+strings.Join(strs[v], " ")], strs[i])
				}
			}
			if MyFlags.A != 0 {
				for i, j := v+1, 0; i <= len(strs)-1; i, j = i+1, j+1 {
					if j > MyFlags.A-1 {
						break
					}
					res["GREP+AFTER:"+strings.Join(strs[v], " ")] = append(res["GREP+AFTER:"+strings.Join(strs[v], " ")], strs[i])
				}
			}
		}
	}
	if MyFlags.C != 0 {
		MyFlags.A, MyFlags.B = MyFlags.C, MyFlags.C
		MyFlags.C = 0
		recMap := grepABC(strs, matchesIdx)
		return recMap
	}
	return res
}
