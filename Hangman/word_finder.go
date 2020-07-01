package main

import (
  "fmt"
  "os"
  b "bufio"
  s "strings"
)

func word_finder(word_to_find string) {

  var letter string
  scanner := b.NewScanner(os.Stdin)
  guess_results := create_empty_guess_results(word_to_find)

  for i := 4; i > 0; i-- {
    fmt.Printf("Enter your guess (%d left): ", i)
    scanner.Scan()

    // check win condition
    if s.ToLower(scanner.Text()) == s.ToLower(word_to_find) && i == 4 {
      fmt.Printf("No one has ever done that, no one has ever done that in the history of hangman!")
      return
    } else if s.ToLower(scanner.Text()) == s.ToLower(word_to_find) {
      fmt.Printf("Mission passed! Respect+")
      return
    }

    if len(scanner.Text()) == 0 {
      fmt.Println("You entered nothing!")
      i++
      continue
    }
    letter = s.Split(scanner.Text(), "")[0]
    fmt.Printf("Your entry was %c\n", letter[0])
    indexes := letter_check(word_to_find, letter)

    filler(indexes, letter, guess_results)
    fmt.Println(guess_results)
  }

  fmt.Printf("Mission failed, we'll get 'em next time.")
  return
}
