package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Question struct {
	q string
	a string
}

func main() {
	csvfile, err := os.Open("quiz.csv")
	if err != nil {
		log.Fatalf("Cannot open the file %v", err)
	}
	
	r := csv.NewReader(csvfile)
	
	fmt.Println("Quiz game")
	fmt.Println("Starting")
	
	recordset, err := r.ReadAll()
	
	if err != nil {
		log.Fatal(err)
	}
	
	questions := parselines(recordset)
	
	correct := 0
	for i, p := range questions {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))
}

func parselines(questions [][]string) []Question {
	ret := make([]Question, len(questions))
	for i, line := range questions {
		ret[i] = Question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
