package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("file", "problems.csv", "csv file with problems in the format of 'question,answer'.")
	timeout := flag.Int("timeout", 30, "time limit in seconds")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("unable to open %s\n", *fileName)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error reading lines from file")
		os.Exit(1)
	}

	problems := parseLines(lines)
	fmt.Println(problems)

	// don't cost the user time while setting up quiz
	timer := time.NewTimer(time.Duration(*timeout) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("\nProblem %d: %s\n", i+1, p.q)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerChannel:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))

	for i, line := range lines {
		result[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return result
}

type problem struct {
	q string
	a string
}
