package main

import (
  "flag"
  "os"
  "fmt"
  "encoding/csv"
  "strings"
)

func main() {
  csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'problem,answer'")
  flag.Parse()

  file, err := os.Open(*csvFilename)
  if err != nil {
    exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
  }

  csvReader := csv.NewReader(file)

  lines, err := csvReader.ReadAll()
  if err != nil {
    exit("Failed to parse the provided csv file.\n")
  }

  problems := parseLines(lines)
  
  problemsCorrect := 0
  for i, problem := range problems {
    fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
    var answer string 
    fmt.Scanf("%s\n", &answer)
    if answer == problem.answer {
      problemsCorrect++
    }
  }
  fmt.Printf("You scored %d out of %d.\n", problemsCorrect, len(problems))
}

func parseLines(lines [][]string) ([]problem) {
  problems := make([]problem, len(lines))
  
  for i, line := range lines {
    problems[i] = problem{
      question: line[0],
      answer: strings.TrimSpace(line[1]),
    }
  }

  return problems
}

type problem struct {
  question string
  answer string
}

func exit(msg string) {
  fmt.Printf(msg)
  os.Exit(1)
}
