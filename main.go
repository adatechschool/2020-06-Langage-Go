package main

import (
  "fmt"
  "bufio"
  "os"
  s "strings"
)

func main() {
  word_finder("fqqsuydf")
}

func word_finder(word_to_find string) {

  var letter, word string
  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < 4; i++{
    fmt.Print("Enter your guess: ")
    scanner.Scan()
    letter = s.Split(scanner.Text(), "")[0]
    fmt.Printf("%c\n", letter[0])
    letter_check(word_to_find, letter)
    word = word + letter
  }

  fmt.Print(word)
  fmt.Print("\n", word_to_find, "\n")
}

// checks for the occurence of the given letter in the word to find
func letter_check(word_to_find, letter string) {
  // array with length of the occurence of the letter to save the indexes
  indexes := make([]int, s.Count(word_to_find, letter))
  // get all indexes of the occurences of "letter"
  placeholder := 0
  for i := 0; i < len(indexes); i++ {
    if i > 0 {
      placeholder = indexes[i-1] + 1
      indexes[i] = s.Index(word_to_find[placeholder:], letter) + placeholder
    } else {
      indexes[i] = s.Index(word_to_find[placeholder:], letter)
    }
  }
  // Result message
  if len(indexes) == 0 {
    fmt.Printf("No occurence of %s in the word.\n", letter, word_to_find)
  } else {
    fmt.Printf("All indexes of %s in the word: %v\n", letter, indexes)
  }
}
