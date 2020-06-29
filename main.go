package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
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
    letter = strings.Split(scanner.Text(), "")[0]
    fmt.Printf("%c\n", letter[0])
    word = word + letter

  }
  fmt.Print(word)
  fmt.Print("\n", word_to_find)
  letter_check(word_to_find, "r")
}

func letter_check(word_to_find, letter string) {
  fmt.Print(strings.Index(word_to_find, letter))
}
