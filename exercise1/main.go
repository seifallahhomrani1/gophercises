package main

import (
  "encoding/csv"
  "fmt"
  "os"
  "log"
  "strconv"
  "io"
  "bufio"
  "strings"
)

func main() {


  var answersCounter,questionsCounter int = 0,0

  csvfile,err := os.Open("tmp_problems.csv")
  if err != nil {
    log.Fatalln("Couldn't open the csv file", err)
  }

  r := csv.NewReader(csvfile)

  for {
    // Reading the quiz questions
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    questionsCounter++ // incrimenting the quiz counter
    // Printing the Quiz
    fmt.Printf("%s = ",record[0])
    // Reading the user's answer
    reader := bufio.NewReader(os.Stdin)
    input,err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("An error occured while reading input, Please try again.", err)
      return
    }

    input = strings.TrimSuffix(input,"\n")

    answer,_ := strconv.ParseInt(input,10,64)
    if err != nil {
      fmt.Println("Error while converting result", err)
      return
    }

    // comparing with the result
    result, _ := strconv.ParseInt(record[1],10,64)

    if answer == result {
      answersCounter++
    }
  }
  fmt.Printf("Final results : %d/%d\n", answersCounter,questionsCounter)
}
