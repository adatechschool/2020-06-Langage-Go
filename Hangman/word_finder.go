package main

import (
  "fmt"
  "os"
  b "bufio"
  s "strings"
)

func word_finder(word_to_find string, guesses int) {
  fmt.Println(word_to_find)

  var letter string
  scanner := b.NewScanner(os.Stdin)
  guess_results := create_empty_guess_results(word_to_find)
  fmt.Println(guess_results)

  for i := guesses; i > 0; i-- {
    fmt.Printf("Enter your guess (%d left): ", i)
    scanner.Scan()

    check_bool := s.ToUpper(scanner.Text()) == s.ToUpper(word_to_find) || s.ToUpper(scanner.Text()) == s.Join(guess_results, "")
    // check win condition
    if check_bool && i == guesses {
      fmt.Printf("No one has ever done that, no one has ever done that in the history of hangman!")
      return
    } else if check_bool {
      fmt.Printf("Mission passed! Respect+")
      return
    }

    if len(scanner.Text()) == 0 {
      fmt.Println("You entered nothing!")
      i++
      continue
    }

    letter = s.ToUpper(s.Split(scanner.Text(), "")[0])
    fmt.Printf("Your entry was %c\n", letter[0])
    indexes := letter_check(word_to_find, letter)
    filler(indexes, letter, guess_results)


    result := s.Join(guess_results, "") // delete later
    fmt.Println("\n[",len(guess_results), len(result),"]") // delete later
  }

  fmt.Printf("Mission failed, we'll get 'em next time.")
  return
}
