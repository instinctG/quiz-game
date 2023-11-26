package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type quiz struct {
	q string
	a string
}

func main() {
	data := make(chan string)

	var correct int
	open := flag.String("csv", "problem.csv", "opens a csv file")
	timer := flag.Int("timer", 20, "set a timer for quiz")
	flag.Parse()
	file, err := os.Open(*open)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := csv.NewReader(file)
	row, err1 := f.ReadAll()
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	problems := parseLines(row)
	timer1 := time.NewTimer(time.Duration(*timer) * time.Second)
	for i, p := range problems {
		var answer string
		fmt.Printf("Question #%v\n%v=", i+1, p.q)
		go func() {

			fmt.Scan(&answer)
			data <- answer

		}()
		select {
		case <-timer1.C:
			fmt.Printf("\nyou`ve scored %v of %v", correct, len(problems))
			return
		case <-data:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("\nyou`ve scored %v of %v", correct, len(problems))
}

func parseLines(row [][]string) []quiz {
	ret := make([]quiz, len(row))
	for i, line := range row {
		ret[i] = quiz{q: line[0],
			a: strings.TrimSpace(line[1])}
	}
	return ret
}
