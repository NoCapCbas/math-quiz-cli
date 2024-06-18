package main

import (
  "flag"
  "os"
  "fmt"
  "encoding/csv"
  "strings"
  "time"
)

func main() {
  csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'problem,answer'")
  timeLimit := flag.Int("limit", 10, "Total time limit for the quiz in seconds")
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

  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
  
  problemsCorrect := 0

problemLoop:
  for i, problem := range problems {
    fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
    answerCh := make(chan string)
    go func(){
      var answer string 
      fmt.Scanf("%s\n", &answer) 
      answerCh <- answer
    }()

    select {
      case <-timer.C:
        fmt.Println()
        fmt.Printf("Time Limit of %d seconds has been reached.\n", *timeLimit)
        break problemLoop
      case answer := <-answerCh:
        if answer == problem.answer {
          problemsCorrect++
        }
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
