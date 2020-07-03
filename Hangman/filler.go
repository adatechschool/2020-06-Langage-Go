package main

import ("fmt")

// place matching letters in the empty array
func filler(indexes []int, letter string, guess_results []string) []string {
  for i := range indexes {
    guess_results[indexes[i]] = letter
  }
  fmt.Println(len(guess_results))
  return guess_results
}
