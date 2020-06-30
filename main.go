package main

import (
  "fmt"
  "bufio"
  "os"
  s "strings"
)

func main() {
  word_finder("florian")
}

func word_finder(word_to_find string) {

  var letter, word string
  scanner := bufio.NewScanner(os.Stdin)
  guess_results := create_empty_guess_results(word_to_find)

  for i := 4; i > 0; i-- {
    fmt.Printf("Enter your guess (%d left): ", i)
    scanner.Scan()
    letter = s.Split(scanner.Text(), "")[0]
    fmt.Printf("Your entry was %c\n", letter[0])
    indexes := letter_check(word_to_find, letter)
    // send indexes to a function that fills out a blank string
    word = word + letter
    filler(indexes, letter, guess_results)
  }
  fmt.Print(word)
  fmt.Print("\n", word_to_find, "\n")
}

// checks for the occurence of the given letter in the word to find
func letter_check(word_to_find, letter string) []int {
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
  return indexes
}

// create the string in which the guessed numbers will be stocked at their respective index
func create_empty_guess_results(word_to_find string) []string {
  empty_results := make([]string, len(word_to_find))
  for i := range empty_results {
    empty_results[i] = "_"
  }
  fmt.Print(empty_results)
  return empty_results
}

func filler(indexes []int, letter string, guess_results []string) {
  for i := range indexes {
    guess_results[indexes[i]] = letter
  }
  // guess_results = s.Replace(guess_results, "_", "t", 4)
  fmt.Print(guess_results)
}
