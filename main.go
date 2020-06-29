package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  word_finder("jean pierre")
}

func word_finder(word_to_find string) {

  var letter, word string
  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < 4; i++{
    fmt.Print("Enter your guess: ")
    scanner.Scan()
    letter = scanner.Text()
    fmt.Printf("%s\n", letter)
    word = word + letter

  }
  fmt.Print(word)
  fmt.Print("\n", word_to_find)
}
